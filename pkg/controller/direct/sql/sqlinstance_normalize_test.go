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

package sql

import (
	"context"
	"testing"

	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sql/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestSQLInstance_NormalizeReferences(t *testing.T) {
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

	t.Run("resolves EncryptionKMSCryptoKeyRef by name", func(t *testing.T) {
		obj := &krm.SQLInstance{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-instance",
				Namespace: "test-ns",
			},
			Spec: krm.SQLInstanceSpec{
				EncryptionKMSCryptoKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
					Name: "test-key",
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		if obj.Spec.EncryptionKMSCryptoKeyRef.External != selfLink {
			t.Errorf("NormalizeReferences() External = %q, want %q", obj.Spec.EncryptionKMSCryptoKeyRef.External, selfLink)
		}

		gcpConfig := InstanceEncryptionKMSCryptoKeyRefKRMToGCP(obj.Spec.EncryptionKMSCryptoKeyRef)
		if gcpConfig == nil || gcpConfig.KmsKeyName != selfLink {
			t.Errorf("InstanceEncryptionKMSCryptoKeyRefKRMToGCP() KmsKeyName = %v, want %q", gcpConfig, selfLink)
		}

		fromGCP := InstanceEncryptionKMSCryptoKeyRefGCPToKRM(gcpConfig)
		if fromGCP == nil || fromGCP.External != selfLink {
			t.Errorf("InstanceEncryptionKMSCryptoKeyRefGCPToKRM() External = %v, want %q", fromGCP, selfLink)
		}
	})

	t.Run("handles external EncryptionKMSCryptoKeyRef directly", func(t *testing.T) {
		externalKey := "projects/test-project/locations/us-central1/keyRings/test-keyring/cryptoKeys/external-key"
		obj := &krm.SQLInstance{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-instance",
				Namespace: "test-ns",
			},
			Spec: krm.SQLInstanceSpec{
				EncryptionKMSCryptoKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
					External: externalKey,
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		if obj.Spec.EncryptionKMSCryptoKeyRef.External != externalKey {
			t.Errorf("NormalizeReferences() External = %q, want %q", obj.Spec.EncryptionKMSCryptoKeyRef.External, externalKey)
		}

		gcpConfig := InstanceEncryptionKMSCryptoKeyRefKRMToGCP(obj.Spec.EncryptionKMSCryptoKeyRef)
		if gcpConfig == nil || gcpConfig.KmsKeyName != externalKey {
			t.Errorf("InstanceEncryptionKMSCryptoKeyRefKRMToGCP() KmsKeyName = %v, want %q", gcpConfig, externalKey)
		}

		fromGCP := InstanceEncryptionKMSCryptoKeyRefGCPToKRM(gcpConfig)
		if fromGCP == nil || fromGCP.External != externalKey {
			t.Errorf("InstanceEncryptionKMSCryptoKeyRefGCPToKRM() External = %v, want %q", fromGCP, externalKey)
		}
	})

	t.Run("handles nil EncryptionKMSCryptoKeyRef", func(t *testing.T) {
		obj := &krm.SQLInstance{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-instance",
				Namespace: "test-ns",
			},
			Spec: krm.SQLInstanceSpec{},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		gcpConfig := InstanceEncryptionKMSCryptoKeyRefKRMToGCP(obj.Spec.EncryptionKMSCryptoKeyRef)
		if gcpConfig != nil {
			t.Errorf("InstanceEncryptionKMSCryptoKeyRefKRMToGCP() = %v, want nil", gcpConfig)
		}
	})
}
