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

package container

import (
	"context"
	"testing"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/container/v1beta1"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestContainerCluster_NormalizeReferences(t *testing.T) {
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

	t.Run("resolves KMSCryptoKeyRef by name in autoProvisioningDefaults", func(t *testing.T) {
		obj := &krm.ContainerCluster{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-cluster",
				Namespace: "test-ns",
			},
			Spec: krm.ContainerClusterSpec{
				ClusterAutoscaling: &krm.ClusterAutoscaling{
					ClusterAutoscaling_AutoProvisioningDefaults: &krm.ClusterAutoscaling_AutoProvisioningDefaults{
						BootDiskKMSKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
							Name: "test-key",
						},
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		got := obj.Spec.ClusterAutoscaling.ClusterAutoscaling_AutoProvisioningDefaults.BootDiskKMSKeyRef.External
		if got != selfLink {
			t.Errorf("NormalizeReferences() External = %q, want %q", got, selfLink)
		}

		mapCtx := &direct.MapContext{}
		protoObj := ClusterAutoscaling_AutoProvisioningDefaults_ToProto(mapCtx, obj.Spec.ClusterAutoscaling.ClusterAutoscaling_AutoProvisioningDefaults)
		if mapCtx.Err() != nil {
			t.Fatalf("ClusterAutoscaling_AutoProvisioningDefaults_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj.GetBootDiskKmsKey() != selfLink {
			t.Errorf("protoObj BootDiskKmsKey = %q, want %q", protoObj.GetBootDiskKmsKey(), selfLink)
		}

		fromProto := ClusterAutoscaling_AutoProvisioningDefaults_FromProto(mapCtx, protoObj)
		if mapCtx.Err() != nil {
			t.Fatalf("ClusterAutoscaling_AutoProvisioningDefaults_FromProto unexpected error: %v", mapCtx.Err())
		}
		if fromProto == nil || fromProto.BootDiskKMSKeyRef == nil || fromProto.BootDiskKMSKeyRef.External != selfLink {
			t.Errorf("fromProto BootDiskKMSKeyRef = %v, want %q", fromProto.BootDiskKMSKeyRef, selfLink)
		}
	})

	t.Run("resolves KMSCryptoKeyRef by name in nodeConfig", func(t *testing.T) {
		obj := &krm.ContainerCluster{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-cluster",
				Namespace: "test-ns",
			},
			Spec: krm.ContainerClusterSpec{
				NodeConfig: &krm.NodeConfig{
					BootDiskKMSCryptoKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
						Name: "test-key",
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		got := obj.Spec.NodeConfig.BootDiskKMSCryptoKeyRef.External
		if got != selfLink {
			t.Errorf("NormalizeReferences() External = %q, want %q", got, selfLink)
		}
	})

	t.Run("handles external KMSCryptoKeyRef directly", func(t *testing.T) {
		externalKey := "projects/test-project/locations/us-central1/keyRings/test-keyring/cryptoKeys/external-key"
		obj := &krm.ContainerCluster{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-cluster",
				Namespace: "test-ns",
			},
			Spec: krm.ContainerClusterSpec{
				ClusterAutoscaling: &krm.ClusterAutoscaling{
					ClusterAutoscaling_AutoProvisioningDefaults: &krm.ClusterAutoscaling_AutoProvisioningDefaults{
						BootDiskKMSKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
							External: externalKey,
						},
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		got := obj.Spec.ClusterAutoscaling.ClusterAutoscaling_AutoProvisioningDefaults.BootDiskKMSKeyRef.External
		if got != externalKey {
			t.Errorf("NormalizeReferences() External = %q, want %q", got, externalKey)
		}
	})

	t.Run("returns error when referenced key is not found", func(t *testing.T) {
		obj := &krm.ContainerCluster{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-cluster",
				Namespace: "test-ns",
			},
			Spec: krm.ContainerClusterSpec{
				ClusterAutoscaling: &krm.ClusterAutoscaling{
					ClusterAutoscaling_AutoProvisioningDefaults: &krm.ClusterAutoscaling_AutoProvisioningDefaults{
						BootDiskKMSKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
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

func TestContainerNodePool_NormalizeReferences(t *testing.T) {
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

	t.Run("resolves KMSCryptoKeyRef by name in nodeConfig", func(t *testing.T) {
		obj := &krm.ContainerNodePool{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-nodepool",
				Namespace: "test-ns",
			},
			Spec: krm.ContainerNodePoolSpec{
				ClusterRef: krm.ContainerClusterRef{
					External: "projects/test-project/locations/us-central1/clusters/test-cluster",
				},
				NodeConfig: &krm.NodePoolNodeConfig{
					BootDiskKMSCryptoKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
						Name: "test-key",
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		got := obj.Spec.NodeConfig.BootDiskKMSCryptoKeyRef.External
		if got != selfLink {
			t.Errorf("NormalizeReferences() External = %q, want %q", got, selfLink)
		}
	})

	t.Run("handles external KMSCryptoKeyRef directly", func(t *testing.T) {
		externalKey := "projects/test-project/locations/us-central1/keyRings/test-keyring/cryptoKeys/external-key"
		obj := &krm.ContainerNodePool{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-nodepool",
				Namespace: "test-ns",
			},
			Spec: krm.ContainerNodePoolSpec{
				ClusterRef: krm.ContainerClusterRef{
					External: "projects/test-project/locations/us-central1/clusters/test-cluster",
				},
				NodeConfig: &krm.NodePoolNodeConfig{
					BootDiskKMSCryptoKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
						External: externalKey,
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		got := obj.Spec.NodeConfig.BootDiskKMSCryptoKeyRef.External
		if got != externalKey {
			t.Errorf("NormalizeReferences() External = %q, want %q", got, externalKey)
		}
	})

	t.Run("returns error when referenced key is not found", func(t *testing.T) {
		obj := &krm.ContainerNodePool{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-nodepool",
				Namespace: "test-ns",
			},
			Spec: krm.ContainerNodePoolSpec{
				ClusterRef: krm.ContainerClusterRef{
					External: "projects/test-project/locations/us-central1/clusters/test-cluster",
				},
				NodeConfig: &krm.NodePoolNodeConfig{
					BootDiskKMSCryptoKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
						Name: "nonexistent-key",
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err == nil {
			t.Fatalf("NormalizeReferences() expected error for nonexistent key, got nil")
		}
	})
}
