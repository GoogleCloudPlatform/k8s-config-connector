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

package gkebackup

import (
	"context"
	"testing"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkebackup/v1alpha1"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestGKEBackupBackupPlan_NormalizeReferences(t *testing.T) {
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
		obj := &krm.GKEBackupBackupPlan{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-backupplan",
				Namespace: "test-ns",
			},
			Spec: krm.GKEBackupBackupPlanSpec{
				BackupConfig: &krm.BackupPlan_BackupConfig{
					EncryptionKey: &krm.EncryptionKey{
						KMSKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
							Name: "test-key",
						},
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		if obj.Spec.BackupConfig.EncryptionKey.KMSKeyRef.External != selfLink {
			t.Errorf("NormalizeReferences() External = %q, want %q", obj.Spec.BackupConfig.EncryptionKey.KMSKeyRef.External, selfLink)
		}

		mapCtx := &direct.MapContext{}
		protoObj := GKEBackupBackupPlanSpec_ToProto(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("GKEBackupBackupPlanSpec_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj.GetBackupConfig().GetEncryptionKey().GetGcpKmsEncryptionKey() != selfLink {
			t.Errorf("protoObj GcpKmsEncryptionKey = %q, want %q", protoObj.GetBackupConfig().GetEncryptionKey().GetGcpKmsEncryptionKey(), selfLink)
		}

		fromProto := GKEBackupBackupPlanSpec_FromProto(mapCtx, protoObj)
		if mapCtx.Err() != nil {
			t.Fatalf("GKEBackupBackupPlanSpec_FromProto unexpected error: %v", mapCtx.Err())
		}
		if fromProto.BackupConfig == nil || fromProto.BackupConfig.EncryptionKey == nil || fromProto.BackupConfig.EncryptionKey.KMSKeyRef == nil || fromProto.BackupConfig.EncryptionKey.KMSKeyRef.External != selfLink {
			t.Errorf("fromProto KMSKeyRef = %v, want external %q", fromProto.BackupConfig.EncryptionKey.KMSKeyRef, selfLink)
		}
	})

	t.Run("handles external KMSKeyRef directly", func(t *testing.T) {
		externalKey := "projects/test-project/locations/us-central1/keyRings/test-keyring/cryptoKeys/external-key"
		obj := &krm.GKEBackupBackupPlan{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-backupplan",
				Namespace: "test-ns",
			},
			Spec: krm.GKEBackupBackupPlanSpec{
				BackupConfig: &krm.BackupPlan_BackupConfig{
					EncryptionKey: &krm.EncryptionKey{
						KMSKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
							External: externalKey,
						},
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		if obj.Spec.BackupConfig.EncryptionKey.KMSKeyRef.External != externalKey {
			t.Errorf("NormalizeReferences() External = %q, want %q", obj.Spec.BackupConfig.EncryptionKey.KMSKeyRef.External, externalKey)
		}

		mapCtx := &direct.MapContext{}
		protoObj := GKEBackupBackupPlanSpec_ToProto(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("GKEBackupBackupPlanSpec_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj.GetBackupConfig().GetEncryptionKey().GetGcpKmsEncryptionKey() != externalKey {
			t.Errorf("protoObj GcpKmsEncryptionKey = %q, want %q", protoObj.GetBackupConfig().GetEncryptionKey().GetGcpKmsEncryptionKey(), externalKey)
		}

		fromProto := GKEBackupBackupPlanSpec_FromProto(mapCtx, protoObj)
		if mapCtx.Err() != nil {
			t.Fatalf("GKEBackupBackupPlanSpec_FromProto unexpected error: %v", mapCtx.Err())
		}
		if fromProto.BackupConfig == nil || fromProto.BackupConfig.EncryptionKey == nil || fromProto.BackupConfig.EncryptionKey.KMSKeyRef == nil || fromProto.BackupConfig.EncryptionKey.KMSKeyRef.External != externalKey {
			t.Errorf("fromProto KMSKeyRef = %v, want external %q", fromProto.BackupConfig.EncryptionKey.KMSKeyRef, externalKey)
		}
	})

	t.Run("handles nil KMSKeyRef", func(t *testing.T) {
		obj := &krm.GKEBackupBackupPlan{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-backupplan",
				Namespace: "test-ns",
			},
			Spec: krm.GKEBackupBackupPlanSpec{
				BackupConfig: &krm.BackupPlan_BackupConfig{},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		mapCtx := &direct.MapContext{}
		protoObj := GKEBackupBackupPlanSpec_ToProto(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("GKEBackupBackupPlanSpec_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj.GetBackupConfig().GetEncryptionKey().GetGcpKmsEncryptionKey() != "" {
			t.Errorf("protoObj GcpKmsEncryptionKey = %q, want empty", protoObj.GetBackupConfig().GetEncryptionKey().GetGcpKmsEncryptionKey())
		}
	})
}
