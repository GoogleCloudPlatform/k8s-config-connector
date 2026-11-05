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

package vertexai

import (
	"context"
	"math/rand"
	"testing"

	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestVertexAIDataset_NormalizeReferences(t *testing.T) {
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

	t.Run("resolves KmsKeyNameRef by name", func(t *testing.T) {
		obj := &krm.VertexAIDataset{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-dataset",
				Namespace: "test-ns",
			},
			Spec: krm.VertexAIDatasetSpec{
				ProjectRef: refs.ProjectRef{
					External: "test-project",
				},
				EncryptionSpec: &krm.DatasetEncryptionSpec{
					KmsKeyNameRef: &kmsv1beta1.KMSCryptoKeyRef{
						Name: "test-key",
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		if obj.Spec.EncryptionSpec.KmsKeyNameRef.External != selfLink {
			t.Errorf("NormalizeReferences() External = %q, want %q", obj.Spec.EncryptionSpec.KmsKeyNameRef.External, selfLink)
		}

		mapCtx := &direct.MapContext{}
		protoObj := DatasetEncryptionSpec_v1beta1_ToProto(mapCtx, obj.Spec.EncryptionSpec)
		if mapCtx.Err() != nil {
			t.Fatalf("DatasetEncryptionSpec_v1beta1_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj.KmsKeyName != selfLink {
			t.Errorf("protoObj KmsKeyName = %q, want %q", protoObj.KmsKeyName, selfLink)
		}

		fromProto := DatasetEncryptionSpec_v1beta1_FromProto(mapCtx, protoObj)
		if mapCtx.Err() != nil {
			t.Fatalf("DatasetEncryptionSpec_v1beta1_FromProto unexpected error: %v", mapCtx.Err())
		}
		if fromProto.KmsKeyNameRef == nil || fromProto.KmsKeyNameRef.External != selfLink {
			t.Errorf("fromProto KmsKeyNameRef.External = %v, want %q", fromProto.KmsKeyNameRef, selfLink)
		}
	})

	t.Run("handles external KmsKeyNameRef directly", func(t *testing.T) {
		externalKey := "projects/test-project/locations/us-central1/keyRings/test-keyring/cryptoKeys/external-key"
		obj := &krm.VertexAIDataset{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-dataset",
				Namespace: "test-ns",
			},
			Spec: krm.VertexAIDatasetSpec{
				ProjectRef: refs.ProjectRef{
					External: "test-project",
				},
				EncryptionSpec: &krm.DatasetEncryptionSpec{
					KmsKeyNameRef: &kmsv1beta1.KMSCryptoKeyRef{
						External: externalKey,
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		if obj.Spec.EncryptionSpec.KmsKeyNameRef.External != externalKey {
			t.Errorf("NormalizeReferences() External = %q, want %q", obj.Spec.EncryptionSpec.KmsKeyNameRef.External, externalKey)
		}

		mapCtx := &direct.MapContext{}
		protoObj := DatasetEncryptionSpec_v1beta1_ToProto(mapCtx, obj.Spec.EncryptionSpec)
		if mapCtx.Err() != nil {
			t.Fatalf("DatasetEncryptionSpec_v1beta1_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj.KmsKeyName != externalKey {
			t.Errorf("protoObj KmsKeyName = %q, want %q", protoObj.KmsKeyName, externalKey)
		}
	})

	t.Run("returns error when referenced key is not found", func(t *testing.T) {
		obj := &krm.VertexAIDataset{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-dataset",
				Namespace: "test-ns",
			},
			Spec: krm.VertexAIDatasetSpec{
				ProjectRef: refs.ProjectRef{
					External: "test-project",
				},
				EncryptionSpec: &krm.DatasetEncryptionSpec{
					KmsKeyNameRef: &kmsv1beta1.KMSCryptoKeyRef{
						Name: "nonexistent-key",
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err == nil {
			t.Fatalf("NormalizeReferences() expected error for nonexistent key, got nil")
		}
	})

	t.Run("handles nil EncryptionSpec gracefully", func(t *testing.T) {
		obj := &krm.VertexAIDataset{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-dataset",
				Namespace: "test-ns",
			},
			Spec: krm.VertexAIDatasetSpec{
				ProjectRef: refs.ProjectRef{
					External: "test-project",
				}},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}
	})

	t.Run("fuzzer roundtrips successfully", func(t *testing.T) {
		fuzzer := vertexAIDatasetFuzzer()
		for i := int64(0); i < 200; i++ {
			seed := rand.Int63()
			fuzzer.FuzzSpec(t, seed)
			fuzzer.FuzzStatus(t, seed)
		}
	})
}
