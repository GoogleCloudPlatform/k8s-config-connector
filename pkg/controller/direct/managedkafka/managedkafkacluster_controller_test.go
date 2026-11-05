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

package managedkafka

import (
	"context"
	"testing"

	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/managedkafka/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestManagedKafkaCluster_NormalizeReferences(t *testing.T) {
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
		obj := &krm.ManagedKafkaCluster{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-cluster",
				Namespace: "test-ns",
			},
			Spec: krm.ManagedKafkaClusterSpec{
				GcpConfig: &krm.GcpConfig{
					KMSKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
						Name: "test-key",
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		if obj.Spec.GcpConfig.KMSKeyRef.External != selfLink {
			t.Errorf("NormalizeReferences() External = %q, want %q", obj.Spec.GcpConfig.KMSKeyRef.External, selfLink)
		}

		mapCtx := &direct.MapContext{}
		protoObj := ManagedKafkaClusterSpec_v1beta1_ToProto(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("ManagedKafkaClusterSpec_v1beta1_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj.GetGcpConfig().GetKmsKey() != selfLink {
			t.Errorf("protoObj KmsKey = %q, want %q", protoObj.GetGcpConfig().GetKmsKey(), selfLink)
		}

		fromProto := ManagedKafkaClusterSpec_v1beta1_FromProto(mapCtx, protoObj)
		if mapCtx.Err() != nil {
			t.Fatalf("ManagedKafkaClusterSpec_v1beta1_FromProto unexpected error: %v", mapCtx.Err())
		}
		if fromProto.GcpConfig == nil || fromProto.GcpConfig.KMSKeyRef == nil || fromProto.GcpConfig.KMSKeyRef.External != selfLink {
			t.Errorf("fromProto KMSKeyRef = %v, want external %q", fromProto.GcpConfig.KMSKeyRef, selfLink)
		}
	})

	t.Run("handles external KMSKeyRef directly", func(t *testing.T) {
		externalKey := "projects/test-project/locations/us-central1/keyRings/test-keyring/cryptoKeys/external-key"
		obj := &krm.ManagedKafkaCluster{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-cluster",
				Namespace: "test-ns",
			},
			Spec: krm.ManagedKafkaClusterSpec{
				GcpConfig: &krm.GcpConfig{
					KMSKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
						External: externalKey,
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		if obj.Spec.GcpConfig.KMSKeyRef.External != externalKey {
			t.Errorf("NormalizeReferences() External = %q, want %q", obj.Spec.GcpConfig.KMSKeyRef.External, externalKey)
		}

		mapCtx := &direct.MapContext{}
		protoObj := ManagedKafkaClusterSpec_v1beta1_ToProto(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("ManagedKafkaClusterSpec_v1beta1_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj.GetGcpConfig().GetKmsKey() != externalKey {
			t.Errorf("protoObj KmsKey = %q, want %q", protoObj.GetGcpConfig().GetKmsKey(), externalKey)
		}

		fromProto := ManagedKafkaClusterSpec_v1beta1_FromProto(mapCtx, protoObj)
		if mapCtx.Err() != nil {
			t.Fatalf("ManagedKafkaClusterSpec_v1beta1_FromProto unexpected error: %v", mapCtx.Err())
		}
		if fromProto.GcpConfig == nil || fromProto.GcpConfig.KMSKeyRef == nil || fromProto.GcpConfig.KMSKeyRef.External != externalKey {
			t.Errorf("fromProto KMSKeyRef = %v, want external %q", fromProto.GcpConfig.KMSKeyRef, externalKey)
		}
	})

	t.Run("handles nil KMSKeyRef", func(t *testing.T) {
		obj := &krm.ManagedKafkaCluster{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-cluster",
				Namespace: "test-ns",
			},
			Spec: krm.ManagedKafkaClusterSpec{
				GcpConfig: &krm.GcpConfig{},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		mapCtx := &direct.MapContext{}
		protoObj := ManagedKafkaClusterSpec_v1beta1_ToProto(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("ManagedKafkaClusterSpec_v1beta1_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj.GetGcpConfig().GetKmsKey() != "" {
			t.Errorf("protoObj KmsKey = %q, want empty", protoObj.GetGcpConfig().GetKmsKey())
		}
	})
}
