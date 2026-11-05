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

package pubsub

import (
	"context"
	"testing"

	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/pubsub/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestPubSubTopic_NormalizeReferences(t *testing.T) {
	ctx := context.Background()

	readyKey := &unstructured.Unstructured{}
	readyKey.SetGroupVersionKind(kmsv1beta1.KMSCryptoKeyGVK)
	readyKey.SetName("test-key")
	readyKey.SetNamespace("test-ns")
	if err := unstructured.SetNestedField(readyKey.Object, "projects/test-project/locations/us/keyRings/test-keyring/cryptoKeys/test-key", "status", "selfLink"); err != nil {
		t.Fatalf("failed to set status.selfLink: %v", err)
	}

	scheme := runtime.NewScheme()
	reader := fake.NewClientBuilder().WithScheme(scheme).WithObjects(readyKey).Build()

	t.Run("resolves KmsKeyRef by name", func(t *testing.T) {
		obj := &krm.PubSubTopic{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-topic",
				Namespace: "test-ns",
			},
			Spec: krm.PubSubTopicSpec{
				KmsKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
					Name: "test-key",
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		want := "projects/test-project/locations/us/keyRings/test-keyring/cryptoKeys/test-key"
		if obj.Spec.KmsKeyRef.External != want {
			t.Errorf("NormalizeReferences() External = %q, want %q", obj.Spec.KmsKeyRef.External, want)
		}

		mapCtx := &direct.MapContext{}
		proto := PubSubTopicSpec_ToProto(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("PubSubTopicSpec_ToProto unexpected error: %v", mapCtx.Err())
		}
		if proto.GetKmsKeyName() != want {
			t.Errorf("proto KmsKeyName = %q, want %q", proto.GetKmsKeyName(), want)
		}
	})

	t.Run("handles external KmsKeyRef directly", func(t *testing.T) {
		externalKey := "projects/test-project/locations/us/keyRings/test-keyring/cryptoKeys/external-key"
		obj := &krm.PubSubTopic{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-topic",
				Namespace: "test-ns",
			},
			Spec: krm.PubSubTopicSpec{
				KmsKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
					External: externalKey,
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		if obj.Spec.KmsKeyRef.External != externalKey {
			t.Errorf("NormalizeReferences() External = %q, want %q", obj.Spec.KmsKeyRef.External, externalKey)
		}

		mapCtx := &direct.MapContext{}
		proto := PubSubTopicSpec_ToProto(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("PubSubTopicSpec_ToProto unexpected error: %v", mapCtx.Err())
		}
		if proto.GetKmsKeyName() != externalKey {
			t.Errorf("proto KmsKeyName = %q, want %q", proto.GetKmsKeyName(), externalKey)
		}
	})
}
