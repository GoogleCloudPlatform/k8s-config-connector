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

package apigee

import (
	"context"
	"testing"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1beta1"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestApigeeInstance_NormalizeReferences(t *testing.T) {
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

	t.Run("resolves DiskEncryptionKMSCryptoKeyRef by name", func(t *testing.T) {
		obj := &krm.ApigeeInstance{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-instance",
				Namespace: "test-ns",
			},
			Spec: krm.ApigeeInstanceSpec{
				DiskEncryptionKMSCryptoKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
					Name: "test-key",
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		if obj.Spec.DiskEncryptionKMSCryptoKeyRef.External != selfLink {
			t.Errorf("NormalizeReferences() External = %q, want %q", obj.Spec.DiskEncryptionKMSCryptoKeyRef.External, selfLink)
		}

		mapCtx := &direct.MapContext{}
		apiObj := ApigeeInstanceSpec_ToAPI(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("ApigeeInstanceSpec_ToAPI unexpected error: %v", mapCtx.Err())
		}
		if apiObj.DiskEncryptionKeyName != selfLink {
			t.Errorf("apiObj DiskEncryptionKeyName = %q, want %q", apiObj.DiskEncryptionKeyName, selfLink)
		}
	})

	t.Run("handles external DiskEncryptionKMSCryptoKeyRef directly", func(t *testing.T) {
		externalKey := "projects/test-project/locations/us-central1/keyRings/test-keyring/cryptoKeys/external-key"
		obj := &krm.ApigeeInstance{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-instance",
				Namespace: "test-ns",
			},
			Spec: krm.ApigeeInstanceSpec{
				DiskEncryptionKMSCryptoKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
					External: externalKey,
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		if obj.Spec.DiskEncryptionKMSCryptoKeyRef.External != externalKey {
			t.Errorf("NormalizeReferences() External = %q, want %q", obj.Spec.DiskEncryptionKMSCryptoKeyRef.External, externalKey)
		}

		mapCtx := &direct.MapContext{}
		apiObj := ApigeeInstanceSpec_ToAPI(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("ApigeeInstanceSpec_ToAPI unexpected error: %v", mapCtx.Err())
		}
		if apiObj.DiskEncryptionKeyName != externalKey {
			t.Errorf("apiObj DiskEncryptionKeyName = %q, want %q", apiObj.DiskEncryptionKeyName, externalKey)
		}

		fromAPI := ApigeeInstanceSpec_FromAPI(mapCtx, apiObj)
		if mapCtx.Err() != nil {
			t.Fatalf("ApigeeInstanceSpec_FromAPI unexpected error: %v", mapCtx.Err())
		}
		if fromAPI.DiskEncryptionKMSCryptoKeyRef == nil || fromAPI.DiskEncryptionKMSCryptoKeyRef.External != externalKey {
			t.Errorf("fromAPI DiskEncryptionKMSCryptoKeyRef.External = %v, want %q", fromAPI.DiskEncryptionKMSCryptoKeyRef, externalKey)
		}
	})

	t.Run("handles nil DiskEncryptionKMSCryptoKeyRef", func(t *testing.T) {
		obj := &krm.ApigeeInstance{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-instance",
				Namespace: "test-ns",
			},
			Spec: krm.ApigeeInstanceSpec{},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		mapCtx := &direct.MapContext{}
		apiObj := ApigeeInstanceSpec_ToAPI(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("ApigeeInstanceSpec_ToAPI unexpected error: %v", mapCtx.Err())
		}
		if apiObj.DiskEncryptionKeyName != "" {
			t.Errorf("apiObj DiskEncryptionKeyName = %q, want empty", apiObj.DiskEncryptionKeyName)
		}
	})
}
