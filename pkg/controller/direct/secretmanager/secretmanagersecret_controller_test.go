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

package secretmanager

import (
	"context"
	"testing"

	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	pubsubv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/pubsub/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/secretmanager/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestSecretManagerSecret_NormalizeReferences(t *testing.T) {
	ctx := context.Background()

	readyKey := &unstructured.Unstructured{}
	readyKey.SetGroupVersionKind(kmsv1beta1.KMSCryptoKeyGVK)
	readyKey.SetName("test-key")
	readyKey.SetNamespace("test-ns")
	selfLink := "projects/test-project/locations/us-central1/keyRings/test-keyring/cryptoKeys/test-key"
	if err := unstructured.SetNestedField(readyKey.Object, selfLink, "status", "externalRef"); err != nil {
		t.Fatalf("failed to set status.externalRef: %v", err)
	}

	readyTopic := &unstructured.Unstructured{}
	readyTopic.SetGroupVersionKind(pubsubv1beta1.PubSubTopicGVK)
	readyTopic.SetName("test-topic")
	readyTopic.SetNamespace("test-ns")
	topicSelfLink := "projects/test-project/topics/test-topic"
	if err := unstructured.SetNestedField(readyTopic.Object, topicSelfLink, "status", "externalRef"); err != nil {
		t.Fatalf("failed to set topic status.externalRef: %v", err)
	}

	scheme := runtime.NewScheme()
	reader := fake.NewClientBuilder().WithScheme(scheme).WithObjects(readyKey, readyTopic).Build()

	t.Run("resolves Automatic CMEK KmsKeyRef by name", func(t *testing.T) {
		obj := &krm.SecretManagerSecret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-secret",
				Namespace: "test-ns",
			},
			Spec: krm.SecretManagerSecretSpec{
				Replication: &krm.Replication{
					LegacyAutomatic: &krm.Replication_Automatic{
						CustomerManagedEncryption: &krm.CustomerManagedEncryption{
							KmsKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
								Name: "test-key",
							},
						},
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		if obj.Spec.Replication.LegacyAutomatic.CustomerManagedEncryption.KmsKeyRef.External != selfLink {
			t.Errorf("NormalizeReferences() External = %q, want %q",
				obj.Spec.Replication.LegacyAutomatic.CustomerManagedEncryption.KmsKeyRef.External, selfLink)
		}

		mapCtx := &direct.MapContext{}
		protoObj := CustomerManagedEncryption_ToProto(mapCtx, obj.Spec.Replication.LegacyAutomatic.CustomerManagedEncryption)
		if mapCtx.Err() != nil {
			t.Fatalf("CustomerManagedEncryption_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj.KmsKeyName != selfLink {
			t.Errorf("protoObj KmsKeyName = %q, want %q", protoObj.KmsKeyName, selfLink)
		}

		fromProto := CustomerManagedEncryption_FromProto(mapCtx, protoObj)
		if mapCtx.Err() != nil {
			t.Fatalf("CustomerManagedEncryption_FromProto unexpected error: %v", mapCtx.Err())
		}
		if fromProto.KmsKeyRef == nil || fromProto.KmsKeyRef.External != selfLink {
			t.Errorf("fromProto KmsKeyRef.External = %v, want %q", fromProto.KmsKeyRef, selfLink)
		}
	})

	t.Run("resolves UserManaged CMEK KmsKeyRef and TopicRef by name", func(t *testing.T) {
		loc := "us-central1"
		obj := &krm.SecretManagerSecret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-secret",
				Namespace: "test-ns",
			},
			Spec: krm.SecretManagerSecretSpec{
				Replication: &krm.Replication{
					UserManaged: &krm.Replication_UserManaged{
						Replicas: []krm.Replication_UserManaged_Replica{
							{
								Location: &loc,
								CustomerManagedEncryption: &krm.CustomerManagedEncryption{
									KmsKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
										Name: "test-key",
									},
								},
							},
						},
					},
				},
				TopicRefs: []*krm.TopicRef{
					{
						PubSubTopicRef: &pubsubv1beta1.PubSubTopicRef{
							Name: "test-topic",
						},
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		replicaKms := obj.Spec.Replication.UserManaged.Replicas[0].CustomerManagedEncryption.KmsKeyRef
		if replicaKms.External != selfLink {
			t.Errorf("NormalizeReferences() Replica KMS External = %q, want %q", replicaKms.External, selfLink)
		}

		topicRef := obj.Spec.TopicRefs[0].PubSubTopicRef
		if topicRef.External != topicSelfLink {
			t.Errorf("NormalizeReferences() TopicRef External = %q, want %q", topicRef.External, topicSelfLink)
		}
	})

	t.Run("handles external KmsKeyRef directly", func(t *testing.T) {
		externalKey := "projects/test-project/locations/us-central1/keyRings/test-keyring/cryptoKeys/external-key"
		obj := &krm.SecretManagerSecret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-secret",
				Namespace: "test-ns",
			},
			Spec: krm.SecretManagerSecretSpec{
				Replication: &krm.Replication{
					LegacyAutomatic: &krm.Replication_Automatic{
						CustomerManagedEncryption: &krm.CustomerManagedEncryption{
							KmsKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
								External: externalKey,
							},
						},
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		if obj.Spec.Replication.LegacyAutomatic.CustomerManagedEncryption.KmsKeyRef.External != externalKey {
			t.Errorf("NormalizeReferences() External = %q, want %q",
				obj.Spec.Replication.LegacyAutomatic.CustomerManagedEncryption.KmsKeyRef.External, externalKey)
		}

		mapCtx := &direct.MapContext{}
		protoObj := CustomerManagedEncryption_ToProto(mapCtx, obj.Spec.Replication.LegacyAutomatic.CustomerManagedEncryption)
		if mapCtx.Err() != nil {
			t.Fatalf("CustomerManagedEncryption_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj.KmsKeyName != externalKey {
			t.Errorf("protoObj KmsKeyName = %q, want %q", protoObj.KmsKeyName, externalKey)
		}
	})
}
