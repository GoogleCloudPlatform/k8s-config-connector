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

package bigtable

import (
	"context"
	"testing"

	krmbigtablev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1alpha1"
	krmbigtablev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestBigtableCluster_NormalizeReferences(t *testing.T) {
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

	t.Run("resolves KMSCryptoKeyRef by name", func(t *testing.T) {
		obj := &krmbigtablev1alpha1.BigtableCluster{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-cluster",
				Namespace: "test-ns",
			},
			Spec: krmbigtablev1alpha1.BigtableClusterSpec{
				EncryptionConfig: &krmbigtablev1alpha1.Cluster_EncryptionConfig{
					KMSKeyNameRef: &kmsv1beta1.KMSCryptoKeyRef{
						Name: "test-key",
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		if obj.Spec.EncryptionConfig.KMSKeyNameRef.External != selfLink {
			t.Errorf("NormalizeReferences() External = %q, want %q", obj.Spec.EncryptionConfig.KMSKeyNameRef.External, selfLink)
		}

		mapCtx := &direct.MapContext{}
		protoObj := BigtableClusterSpec_v1alpha1_ToProto(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("BigtableClusterSpec_v1alpha1_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj.GetEncryptionConfig().GetKmsKeyName() != selfLink {
			t.Errorf("protoObj KmsKeyName = %q, want %q", protoObj.GetEncryptionConfig().GetKmsKeyName(), selfLink)
		}

		fromProto := BigtableClusterSpec_v1alpha1_FromProto(mapCtx, protoObj)
		if mapCtx.Err() != nil {
			t.Fatalf("BigtableClusterSpec_v1alpha1_FromProto unexpected error: %v", mapCtx.Err())
		}
		if fromProto.EncryptionConfig == nil || fromProto.EncryptionConfig.KMSKeyNameRef == nil || fromProto.EncryptionConfig.KMSKeyNameRef.External != selfLink {
			t.Errorf("fromProto KMSKeyNameRef.External = %v, want %q", fromProto.EncryptionConfig.KMSKeyNameRef, selfLink)
		}
	})

	t.Run("handles external KMSCryptoKeyRef directly", func(t *testing.T) {
		externalKey := "projects/test-project/locations/us-central1/keyRings/test-keyring/cryptoKeys/external-key"
		obj := &krmbigtablev1alpha1.BigtableCluster{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-cluster",
				Namespace: "test-ns",
			},
			Spec: krmbigtablev1alpha1.BigtableClusterSpec{
				EncryptionConfig: &krmbigtablev1alpha1.Cluster_EncryptionConfig{
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
		protoObj := BigtableClusterSpec_v1alpha1_ToProto(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("BigtableClusterSpec_v1alpha1_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj.GetEncryptionConfig().GetKmsKeyName() != externalKey {
			t.Errorf("protoObj KmsKeyName = %q, want %q", protoObj.GetEncryptionConfig().GetKmsKeyName(), externalKey)
		}
	})

	t.Run("returns error when referenced key is not found", func(t *testing.T) {
		obj := &krmbigtablev1alpha1.BigtableCluster{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-cluster",
				Namespace: "test-ns",
			},
			Spec: krmbigtablev1alpha1.BigtableClusterSpec{
				EncryptionConfig: &krmbigtablev1alpha1.Cluster_EncryptionConfig{
					KMSKeyNameRef: &kmsv1beta1.KMSCryptoKeyRef{
						Name: "nonexistent-key",
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err == nil {
			t.Fatalf("NormalizeReferences() expected error for nonexistent key, got nil")
		}
	})

	t.Run("handles nil KMSKeyNameRef gracefully", func(t *testing.T) {
		obj := &krmbigtablev1alpha1.BigtableCluster{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-cluster",
				Namespace: "test-ns",
			},
			Spec: krmbigtablev1alpha1.BigtableClusterSpec{
				EncryptionConfig: &krmbigtablev1alpha1.Cluster_EncryptionConfig{},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}
	})
}

func TestBigtableInstance_NormalizeReferences(t *testing.T) {
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

	t.Run("resolves KMSCryptoKeyRef by name", func(t *testing.T) {
		obj := &krmbigtablev1beta1.BigtableInstance{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-instance",
				Namespace: "test-ns",
			},
			Spec: krmbigtablev1beta1.BigtableInstanceSpec{
				Cluster: []krmbigtablev1beta1.InstanceCluster{
					{
						ClusterId: "cluster-1",
						Zone:      "us-central1-a",
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

		if obj.Spec.Cluster[0].KmsKeyRef.External != selfLink {
			t.Errorf("NormalizeReferences() External = %q, want %q", obj.Spec.Cluster[0].KmsKeyRef.External, selfLink)
		}
	})

	t.Run("handles external KMSCryptoKeyRef directly", func(t *testing.T) {
		externalKey := "projects/test-project/locations/us-central1/keyRings/test-keyring/cryptoKeys/external-key"
		obj := &krmbigtablev1beta1.BigtableInstance{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-instance",
				Namespace: "test-ns",
			},
			Spec: krmbigtablev1beta1.BigtableInstanceSpec{
				Cluster: []krmbigtablev1beta1.InstanceCluster{
					{
						ClusterId: "cluster-1",
						Zone:      "us-central1-a",
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

		if obj.Spec.Cluster[0].KmsKeyRef.External != externalKey {
			t.Errorf("NormalizeReferences() External = %q, want %q", obj.Spec.Cluster[0].KmsKeyRef.External, externalKey)
		}
	})

	t.Run("returns error when referenced key is not found", func(t *testing.T) {
		obj := &krmbigtablev1beta1.BigtableInstance{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-instance",
				Namespace: "test-ns",
			},
			Spec: krmbigtablev1beta1.BigtableInstanceSpec{
				Cluster: []krmbigtablev1beta1.InstanceCluster{
					{
						ClusterId: "cluster-1",
						Zone:      "us-central1-a",
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
}
