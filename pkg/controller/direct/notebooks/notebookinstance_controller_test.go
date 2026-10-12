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

package notebooks

import (
	"context"
	"math/rand"
	"testing"

	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/notebooks/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestNotebookInstance_NormalizeReferences(t *testing.T) {
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
		obj := &krm.NotebookInstance{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-instance",
				Namespace: "test-ns",
			},
			Spec: krm.NotebookInstanceSpec{
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
		protoObj := NotebookInstanceSpec_v1beta1_ToProto(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("NotebookInstanceSpec_v1beta1_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj.GetKmsKey() != selfLink {
			t.Errorf("protoObj KmsKey = %q, want %q", protoObj.GetKmsKey(), selfLink)
		}

		fromProto := NotebookInstanceSpec_v1beta1_FromProto(mapCtx, protoObj)
		if mapCtx.Err() != nil {
			t.Fatalf("NotebookInstanceSpec_v1beta1_FromProto unexpected error: %v", mapCtx.Err())
		}
		if fromProto.KMSKeyRef == nil || fromProto.KMSKeyRef.External != selfLink {
			t.Errorf("fromProto KMSKeyRef.External = %v, want %q", fromProto.KMSKeyRef, selfLink)
		}
	})

	t.Run("handles external KMSKeyRef directly", func(t *testing.T) {
		externalKey := "projects/test-project/locations/us-central1/keyRings/test-keyring/cryptoKeys/external-key"
		obj := &krm.NotebookInstance{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-instance",
				Namespace: "test-ns",
			},
			Spec: krm.NotebookInstanceSpec{
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
		protoObj := NotebookInstanceSpec_v1beta1_ToProto(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("NotebookInstanceSpec_v1beta1_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj.GetKmsKey() != externalKey {
			t.Errorf("protoObj KmsKey = %q, want %q", protoObj.GetKmsKey(), externalKey)
		}
	})

	t.Run("returns error when referenced key is not found", func(t *testing.T) {
		obj := &krm.NotebookInstance{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-instance",
				Namespace: "test-ns",
			},
			Spec: krm.NotebookInstanceSpec{
				KMSKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
					Name: "nonexistent-key",
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err == nil {
			t.Fatalf("NormalizeReferences() expected error for nonexistent key, got nil")
		}
	})

	t.Run("handles nil KMSKeyRef gracefully", func(t *testing.T) {
		obj := &krm.NotebookInstance{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-instance",
				Namespace: "test-ns",
			},
			Spec: krm.NotebookInstanceSpec{},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}
	})

	t.Run("fuzzer roundtrips successfully", func(t *testing.T) {
		fz := notebookInstanceFuzzer()
		for i := int64(0); i < 200; i++ {
			seed := rand.Int63()
			fz.FuzzSpec(t, seed)
			fz.FuzzStatus(t, seed)
		}
	})
}
