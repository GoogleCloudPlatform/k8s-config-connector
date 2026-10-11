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

package spanner

import (
	"context"
	"testing"

	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/spanner/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestSpannerBackupSchedule_NormalizeReferences(t *testing.T) {
	ctx := context.Background()

	key1 := &unstructured.Unstructured{}
	key1.SetGroupVersionKind(kmsv1beta1.KMSCryptoKeyGVK)
	key1.SetName("test-key-1")
	key1.SetNamespace("test-ns")
	selfLink1 := "projects/test-project/locations/us-central1/keyRings/test-keyring/cryptoKeys/test-key-1"
	if err := unstructured.SetNestedField(key1.Object, selfLink1, "status", "externalRef"); err != nil {
		t.Fatalf("failed to set status.externalRef: %v", err)
	}
	if err := unstructured.SetNestedField(key1.Object, selfLink1, "status", "selfLink"); err != nil {
		t.Fatalf("failed to set status.selfLink: %v", err)
	}

	key2 := &unstructured.Unstructured{}
	key2.SetGroupVersionKind(kmsv1beta1.KMSCryptoKeyGVK)
	key2.SetName("test-key-2")
	key2.SetNamespace("test-ns")
	selfLink2 := "projects/test-project/locations/us-east1/keyRings/test-keyring/cryptoKeys/test-key-2"
	if err := unstructured.SetNestedField(key2.Object, selfLink2, "status", "externalRef"); err != nil {
		t.Fatalf("failed to set status.externalRef: %v", err)
	}
	if err := unstructured.SetNestedField(key2.Object, selfLink2, "status", "selfLink"); err != nil {
		t.Fatalf("failed to set status.selfLink: %v", err)
	}

	scheme := runtime.NewScheme()
	reader := fake.NewClientBuilder().WithScheme(scheme).WithObjects(key1, key2).Build()

	t.Run("resolves KMSKeyRef and KMSKeyRefs by name", func(t *testing.T) {
		obj := &krm.SpannerBackupSchedule{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-schedule",
				Namespace: "test-ns",
			},
			Spec: krm.SpannerBackupScheduleSpec{
				EncryptionConfig: &krm.CreateBackupEncryptionConfig{
					KMSKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
						Name: "test-key-1",
					},
					KMSKeyRefs: []*kmsv1beta1.KMSCryptoKeyRef{
						{Name: "test-key-1"},
						{Name: "test-key-2"},
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		if obj.Spec.EncryptionConfig.KMSKeyRef.External != selfLink1 {
			t.Errorf("NormalizeReferences() KMSKeyRef.External = %q, want %q", obj.Spec.EncryptionConfig.KMSKeyRef.External, selfLink1)
		}
		if len(obj.Spec.EncryptionConfig.KMSKeyRefs) != 2 {
			t.Fatalf("len(KMSKeyRefs) = %d, want 2", len(obj.Spec.EncryptionConfig.KMSKeyRefs))
		}
		if obj.Spec.EncryptionConfig.KMSKeyRefs[0].External != selfLink1 {
			t.Errorf("NormalizeReferences() KMSKeyRefs[0].External = %q, want %q", obj.Spec.EncryptionConfig.KMSKeyRefs[0].External, selfLink1)
		}
		if obj.Spec.EncryptionConfig.KMSKeyRefs[1].External != selfLink2 {
			t.Errorf("NormalizeReferences() KMSKeyRefs[1].External = %q, want %q", obj.Spec.EncryptionConfig.KMSKeyRefs[1].External, selfLink2)
		}

		mapCtx := &direct.MapContext{}
		protoObj := SpannerBackupScheduleSpec_ToProto(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("SpannerBackupScheduleSpec_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj.EncryptionConfig.KmsKeyName != selfLink1 {
			t.Errorf("protoObj KmsKeyName = %q, want %q", protoObj.EncryptionConfig.KmsKeyName, selfLink1)
		}
		if len(protoObj.EncryptionConfig.KmsKeyNames) != 2 || protoObj.EncryptionConfig.KmsKeyNames[0] != selfLink1 || protoObj.EncryptionConfig.KmsKeyNames[1] != selfLink2 {
			t.Errorf("protoObj KmsKeyNames = %v, want [%q, %q]", protoObj.EncryptionConfig.KmsKeyNames, selfLink1, selfLink2)
		}

		fromProto := CreateBackupEncryptionConfig_FromProto(mapCtx, protoObj.EncryptionConfig)
		if mapCtx.Err() != nil {
			t.Fatalf("CreateBackupEncryptionConfig_FromProto unexpected error: %v", mapCtx.Err())
		}
		if fromProto.KMSKeyRef == nil || fromProto.KMSKeyRef.External != selfLink1 {
			t.Errorf("fromProto KMSKeyRef.External = %v, want %q", fromProto.KMSKeyRef, selfLink1)
		}
		if len(fromProto.KMSKeyRefs) != 2 || fromProto.KMSKeyRefs[0].External != selfLink1 || fromProto.KMSKeyRefs[1].External != selfLink2 {
			t.Errorf("fromProto KMSKeyRefs = %v, want [%q, %q]", fromProto.KMSKeyRefs, selfLink1, selfLink2)
		}
	})

	t.Run("handles external KMSKeyRef and KMSKeyRefs directly", func(t *testing.T) {
		externalKey1 := "projects/test-project/locations/us-central1/keyRings/test-keyring/cryptoKeys/ext-key-1"
		externalKey2 := "projects/test-project/locations/us-east1/keyRings/test-keyring/cryptoKeys/ext-key-2"
		obj := &krm.SpannerBackupSchedule{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-schedule",
				Namespace: "test-ns",
			},
			Spec: krm.SpannerBackupScheduleSpec{
				EncryptionConfig: &krm.CreateBackupEncryptionConfig{
					KMSKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
						External: externalKey1,
					},
					KMSKeyRefs: []*kmsv1beta1.KMSCryptoKeyRef{
						{External: externalKey1},
						{External: externalKey2},
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		if obj.Spec.EncryptionConfig.KMSKeyRef.External != externalKey1 {
			t.Errorf("NormalizeReferences() KMSKeyRef.External = %q, want %q", obj.Spec.EncryptionConfig.KMSKeyRef.External, externalKey1)
		}
		if obj.Spec.EncryptionConfig.KMSKeyRefs[0].External != externalKey1 {
			t.Errorf("NormalizeReferences() KMSKeyRefs[0].External = %q, want %q", obj.Spec.EncryptionConfig.KMSKeyRefs[0].External, externalKey1)
		}
		if obj.Spec.EncryptionConfig.KMSKeyRefs[1].External != externalKey2 {
			t.Errorf("NormalizeReferences() KMSKeyRefs[1].External = %q, want %q", obj.Spec.EncryptionConfig.KMSKeyRefs[1].External, externalKey2)
		}

		mapCtx := &direct.MapContext{}
		protoObj := SpannerBackupScheduleSpec_ToProto(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("SpannerBackupScheduleSpec_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj.EncryptionConfig.KmsKeyName != externalKey1 {
			t.Errorf("protoObj KmsKeyName = %q, want %q", protoObj.EncryptionConfig.KmsKeyName, externalKey1)
		}
		if len(protoObj.EncryptionConfig.KmsKeyNames) != 2 || protoObj.EncryptionConfig.KmsKeyNames[0] != externalKey1 || protoObj.EncryptionConfig.KmsKeyNames[1] != externalKey2 {
			t.Errorf("protoObj KmsKeyNames = %v, want [%q, %q]", protoObj.EncryptionConfig.KmsKeyNames, externalKey1, externalKey2)
		}
	})

	t.Run("omits KmsKeyName and KmsKeyNames when EncryptionType is GOOGLE_DEFAULT_ENCRYPTION", func(t *testing.T) {
		encType := "GOOGLE_DEFAULT_ENCRYPTION"
		externalKey := "projects/test-project/locations/us-central1/keyRings/test-keyring/cryptoKeys/ext-key-1"
		obj := &krm.SpannerBackupSchedule{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-schedule",
				Namespace: "test-ns",
			},
			Spec: krm.SpannerBackupScheduleSpec{
				EncryptionConfig: &krm.CreateBackupEncryptionConfig{
					EncryptionType: &encType,
					KMSKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
						External: externalKey,
					},
					KMSKeyRefs: []*kmsv1beta1.KMSCryptoKeyRef{
						{External: externalKey},
					},
				},
			},
		}

		mapCtx := &direct.MapContext{}
		protoObj := SpannerBackupScheduleSpec_ToProto(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("SpannerBackupScheduleSpec_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj.EncryptionConfig == nil {
			t.Fatalf("protoObj.EncryptionConfig is nil")
		}
		if protoObj.EncryptionConfig.KmsKeyName != "" {
			t.Errorf("protoObj KmsKeyName = %q, want empty", protoObj.EncryptionConfig.KmsKeyName)
		}
		if len(protoObj.EncryptionConfig.KmsKeyNames) != 0 {
			t.Errorf("protoObj KmsKeyNames = %v, want empty", protoObj.EncryptionConfig.KmsKeyNames)
		}
	})

	t.Run("handles nil EncryptionConfig", func(t *testing.T) {
		obj := &krm.SpannerBackupSchedule{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-schedule",
				Namespace: "test-ns",
			},
			Spec: krm.SpannerBackupScheduleSpec{},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		mapCtx := &direct.MapContext{}
		protoObj := SpannerBackupScheduleSpec_ToProto(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("SpannerBackupScheduleSpec_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj.EncryptionConfig != nil {
			t.Errorf("protoObj.EncryptionConfig = %v, want nil", protoObj.EncryptionConfig)
		}
	})

	t.Run("fuzzer roundtrips successfully", func(t *testing.T) {
		f := spannerBackupScheduleFuzzer()
		for seed := int64(0); seed < 200; seed++ {
			f.FuzzSpec(t, seed)
			f.FuzzStatus(t, seed)
		}
	})
}
