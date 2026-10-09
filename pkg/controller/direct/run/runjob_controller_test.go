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

package run

import (
	"context"
	"testing"

	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/run/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestRunJob_NormalizeReferences(t *testing.T) {
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

	t.Run("resolves EncryptionKeyRef by name", func(t *testing.T) {
		obj := &krm.RunJob{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-job",
				Namespace: "test-ns",
			},
			Spec: krm.RunJobSpec{
				Template: &krm.ExecutionTemplate{
					Template: &krm.TaskTemplate{
						EncryptionKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
							Name: "test-key",
						},
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		if obj.Spec.Template.Template.EncryptionKeyRef.External != selfLink {
			t.Errorf("NormalizeReferences() External = %q, want %q", obj.Spec.Template.Template.EncryptionKeyRef.External, selfLink)
		}

		if err := ResolveRunJobRefs(ctx, reader, obj); err != nil {
			t.Fatalf("ResolveRunJobRefs() unexpected error: %v", err)
		}

		mapCtx := &direct.MapContext{}
		protoObj := TaskTemplate_v1beta1_ToProto(mapCtx, obj.Spec.Template.Template)
		if mapCtx.Err() != nil {
			t.Fatalf("TaskTemplate_v1beta1_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj.EncryptionKey != selfLink {
			t.Errorf("protoObj EncryptionKey = %q, want %q", protoObj.EncryptionKey, selfLink)
		}

		fromProto := TaskTemplate_v1beta1_FromProto(mapCtx, protoObj)
		if mapCtx.Err() != nil {
			t.Fatalf("TaskTemplate_v1beta1_FromProto unexpected error: %v", mapCtx.Err())
		}
		if fromProto.EncryptionKeyRef == nil || fromProto.EncryptionKeyRef.External != selfLink {
			t.Errorf("fromProto EncryptionKeyRef.External = %v, want %q", fromProto.EncryptionKeyRef, selfLink)
		}
	})

	t.Run("handles external EncryptionKeyRef directly", func(t *testing.T) {
		externalKey := "projects/test-project/locations/us-central1/keyRings/test-keyring/cryptoKeys/external-key"
		obj := &krm.RunJob{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-job",
				Namespace: "test-ns",
			},
			Spec: krm.RunJobSpec{
				Template: &krm.ExecutionTemplate{
					Template: &krm.TaskTemplate{
						EncryptionKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
							External: externalKey,
						},
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		if obj.Spec.Template.Template.EncryptionKeyRef.External != externalKey {
			t.Errorf("NormalizeReferences() External = %q, want %q", obj.Spec.Template.Template.EncryptionKeyRef.External, externalKey)
		}

		if err := ResolveRunJobRefs(ctx, reader, obj); err != nil {
			t.Fatalf("ResolveRunJobRefs() unexpected error: %v", err)
		}

		mapCtx := &direct.MapContext{}
		protoObj := TaskTemplate_v1beta1_ToProto(mapCtx, obj.Spec.Template.Template)
		if mapCtx.Err() != nil {
			t.Fatalf("TaskTemplate_v1beta1_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj.EncryptionKey != externalKey {
			t.Errorf("protoObj EncryptionKey = %q, want %q", protoObj.EncryptionKey, externalKey)
		}
	})
}
