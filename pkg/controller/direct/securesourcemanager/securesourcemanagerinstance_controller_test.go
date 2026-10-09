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

package securesourcemanager

import (
	"context"
	"testing"

	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securesourcemanager/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestSecureSourceManagerInstance_NormalizeReferences(t *testing.T) {
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

	t.Run("resolves KMSKeyRef by name", func(t *testing.T) {
		obj := &krm.SecureSourceManagerInstance{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-instance",
				Namespace: "test-ns",
			},
			Spec: krm.SecureSourceManagerInstanceSpec{
				KMSKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
					Name: "test-key",
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		if obj.Spec.KMSKeyRef.External != selfLink {
			t.Errorf("NormalizeReferences() External = %q, want %q", obj.Spec.KMSKeyRef.External, selfLink)
		}

		mapCtx := &direct.MapContext{}
		protoObj := SecureSourceManagerInstanceSpec_ToProto(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("SecureSourceManagerInstanceSpec_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj.KmsKey != selfLink {
			t.Errorf("protoObj KmsKey = %q, want %q", protoObj.KmsKey, selfLink)
		}

		fromProto := SecureSourceManagerInstanceSpec_FromProto(mapCtx, protoObj)
		if mapCtx.Err() != nil {
			t.Fatalf("SecureSourceManagerInstanceSpec_FromProto unexpected error: %v", mapCtx.Err())
		}
		if fromProto.KMSKeyRef == nil || fromProto.KMSKeyRef.External != selfLink {
			t.Errorf("fromProto KMSKeyRef.External = %v, want %q", fromProto.KMSKeyRef, selfLink)
		}
	})

	t.Run("handles external KMSKeyRef directly", func(t *testing.T) {
		externalKey := "projects/test-project/locations/us-central1/keyRings/test-keyring/cryptoKeys/external-key"
		obj := &krm.SecureSourceManagerInstance{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-instance",
				Namespace: "test-ns",
			},
			Spec: krm.SecureSourceManagerInstanceSpec{
				KMSKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
					External: externalKey,
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		if obj.Spec.KMSKeyRef.External != externalKey {
			t.Errorf("NormalizeReferences() External = %q, want %q", obj.Spec.KMSKeyRef.External, externalKey)
		}

		mapCtx := &direct.MapContext{}
		protoObj := SecureSourceManagerInstanceSpec_ToProto(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("SecureSourceManagerInstanceSpec_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj.KmsKey != externalKey {
			t.Errorf("protoObj KmsKey = %q, want %q", protoObj.KmsKey, externalKey)
		}

		fromProto := SecureSourceManagerInstanceSpec_FromProto(mapCtx, protoObj)
		if mapCtx.Err() != nil {
			t.Fatalf("SecureSourceManagerInstanceSpec_FromProto unexpected error: %v", mapCtx.Err())
		}
		if fromProto.KMSKeyRef == nil || fromProto.KMSKeyRef.External != externalKey {
			t.Errorf("fromProto KMSKeyRef.External = %v, want %q", fromProto.KMSKeyRef, externalKey)
		}
	})

	t.Run("handles nil KMSKeyRef", func(t *testing.T) {
		obj := &krm.SecureSourceManagerInstance{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-instance",
				Namespace: "test-ns",
			},
			Spec: krm.SecureSourceManagerInstanceSpec{},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		mapCtx := &direct.MapContext{}
		protoObj := SecureSourceManagerInstanceSpec_ToProto(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("SecureSourceManagerInstanceSpec_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj.KmsKey != "" {
			t.Errorf("protoObj KmsKey = %q, want empty", protoObj.KmsKey)
		}
	})
}
