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

package alloydb

import (
	"context"
	"testing"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/alloydb/v1beta1"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestAlloyDBCluster_NormalizeReferences(t *testing.T) {
	ctx := context.Background()

	readyKey := &unstructured.Unstructured{}
	readyKey.SetGroupVersionKind(kmsv1beta1.KMSCryptoKeyGVK)
	readyKey.SetName("test-key")
	readyKey.SetNamespace("test-ns")
	selfLink := "projects/test-project/locations/us-central1/keyRings/test-keyring/cryptoKeys/test-key"
	if err := unstructured.SetNestedField(readyKey.Object, selfLink, "status", "externalRef"); err != nil {
		t.Fatalf("failed to set status.externalRef: %v", err)
	}
	if err := unstructured.SetNestedField(readyKey.Object, selfLink, "status", "selfLink"); err != nil {
		t.Fatalf("failed to set status.selfLink: %v", err)
	}

	scheme := runtime.NewScheme()
	reader := fake.NewClientBuilder().WithScheme(scheme).WithObjects(readyKey).Build()

	t.Run("resolves KMSKeyNameRef by name across cluster, automated backup, and continuous backup", func(t *testing.T) {
		obj := &krm.AlloyDBCluster{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-cluster",
				Namespace: "test-ns",
			},
			Spec: krm.AlloyDBClusterSpec{
				EncryptionConfig: &krm.EncryptionConfig{
					KMSKeyNameRef: &kmsv1beta1.KMSCryptoKeyRef{
						Name: "test-key",
					},
				},
				AutomatedBackupPolicy: &krm.AutomatedBackupPolicy{
					EncryptionConfig: &krm.EncryptionConfig{
						KMSKeyNameRef: &kmsv1beta1.KMSCryptoKeyRef{
							Name: "test-key",
						},
					},
				},
				ContinuousBackupConfig: &krm.ContinuousBackupConfig{
					EncryptionConfig: &krm.EncryptionConfig{
						KMSKeyNameRef: &kmsv1beta1.KMSCryptoKeyRef{
							Name: "test-key",
						},
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		if obj.Spec.EncryptionConfig.KMSKeyNameRef.External != selfLink {
			t.Errorf("NormalizeReferences() cluster EncryptionConfig External = %q, want %q", obj.Spec.EncryptionConfig.KMSKeyNameRef.External, selfLink)
		}
		if obj.Spec.AutomatedBackupPolicy.EncryptionConfig.KMSKeyNameRef.External != selfLink {
			t.Errorf("NormalizeReferences() AutomatedBackupPolicy EncryptionConfig External = %q, want %q", obj.Spec.AutomatedBackupPolicy.EncryptionConfig.KMSKeyNameRef.External, selfLink)
		}
		if obj.Spec.ContinuousBackupConfig.EncryptionConfig.KMSKeyNameRef.External != selfLink {
			t.Errorf("NormalizeReferences() ContinuousBackupConfig EncryptionConfig External = %q, want %q", obj.Spec.ContinuousBackupConfig.EncryptionConfig.KMSKeyNameRef.External, selfLink)
		}

		mapCtx := &direct.MapContext{}
		protoObj := AlloyDBClusterSpec_ToProto(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("AlloyDBClusterSpec_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj.GetEncryptionConfig().GetKmsKeyName() != selfLink {
			t.Errorf("proto cluster encryptionConfig kms_key_name = %q, want %q", protoObj.GetEncryptionConfig().GetKmsKeyName(), selfLink)
		}
		if protoObj.GetAutomatedBackupPolicy().GetEncryptionConfig().GetKmsKeyName() != selfLink {
			t.Errorf("proto automatedBackupPolicy encryptionConfig kms_key_name = %q, want %q", protoObj.GetAutomatedBackupPolicy().GetEncryptionConfig().GetKmsKeyName(), selfLink)
		}
		if protoObj.GetContinuousBackupConfig().GetEncryptionConfig().GetKmsKeyName() != selfLink {
			t.Errorf("proto continuousBackupConfig encryptionConfig kms_key_name = %q, want %q", protoObj.GetContinuousBackupConfig().GetEncryptionConfig().GetKmsKeyName(), selfLink)
		}

		fromProto := AlloyDBClusterSpec_FromProto(mapCtx, protoObj)
		if mapCtx.Err() != nil {
			t.Fatalf("AlloyDBClusterSpec_FromProto unexpected error: %v", mapCtx.Err())
		}
		if fromProto.EncryptionConfig == nil || fromProto.EncryptionConfig.KMSKeyNameRef == nil || fromProto.EncryptionConfig.KMSKeyNameRef.External != selfLink {
			t.Errorf("fromProto cluster EncryptionConfig KMSKeyNameRef = %v, want %q", fromProto.EncryptionConfig, selfLink)
		}
		if fromProto.AutomatedBackupPolicy == nil || fromProto.AutomatedBackupPolicy.EncryptionConfig == nil || fromProto.AutomatedBackupPolicy.EncryptionConfig.KMSKeyNameRef.External != selfLink {
			t.Errorf("fromProto automatedBackupPolicy EncryptionConfig KMSKeyNameRef = %v, want %q", fromProto.AutomatedBackupPolicy, selfLink)
		}
		if fromProto.ContinuousBackupConfig == nil || fromProto.ContinuousBackupConfig.EncryptionConfig == nil || fromProto.ContinuousBackupConfig.EncryptionConfig.KMSKeyNameRef.External != selfLink {
			t.Errorf("fromProto continuousBackupConfig EncryptionConfig KMSKeyNameRef = %v, want %q", fromProto.ContinuousBackupConfig, selfLink)
		}
	})

	t.Run("handles external KMSKeyNameRef directly", func(t *testing.T) {
		externalKey := "projects/test-project/locations/us-central1/keyRings/test-keyring/cryptoKeys/external-key"
		obj := &krm.AlloyDBCluster{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-cluster",
				Namespace: "test-ns",
			},
			Spec: krm.AlloyDBClusterSpec{
				EncryptionConfig: &krm.EncryptionConfig{
					KMSKeyNameRef: &kmsv1beta1.KMSCryptoKeyRef{
						External: externalKey,
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		if obj.Spec.EncryptionConfig.KMSKeyNameRef.External != externalKey {
			t.Errorf("NormalizeReferences() External = %q, want %q", obj.Spec.EncryptionConfig.KMSKeyNameRef.External, externalKey)
		}

		mapCtx := &direct.MapContext{}
		protoObj := AlloyDBClusterSpec_ToProto(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("AlloyDBClusterSpec_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj.GetEncryptionConfig().GetKmsKeyName() != externalKey {
			t.Errorf("proto cluster encryptionConfig kms_key_name = %q, want %q", protoObj.GetEncryptionConfig().GetKmsKeyName(), externalKey)
		}
	})

	t.Run("handles nil EncryptionConfig gracefully", func(t *testing.T) {
		obj := &krm.AlloyDBCluster{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-cluster",
				Namespace: "test-ns",
			},
			Spec: krm.AlloyDBClusterSpec{},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		mapCtx := &direct.MapContext{}
		protoObj := AlloyDBClusterSpec_ToProto(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("AlloyDBClusterSpec_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj.GetEncryptionConfig() != nil {
			t.Errorf("proto cluster encryptionConfig = %v, want nil", protoObj.GetEncryptionConfig())
		}
	})
}
