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

package dataproc

import (
	"context"
	"testing"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataproc/v1beta1"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestDataprocCluster_NormalizeReferences(t *testing.T) {
	ctx := context.Background()

	readyKey := &unstructured.Unstructured{}
	readyKey.SetGroupVersionKind(kmsv1beta1.KMSCryptoKeyGVK)
	readyKey.SetName("test-kms-key")
	readyKey.SetNamespace("test-ns")
	selfLink := "projects/test-project/locations/us-central1/keyRings/test-keyring/cryptoKeys/test-kms-key"
	if err := unstructured.SetNestedField(readyKey.Object, selfLink, "status", "selfLink"); err != nil {
		t.Fatalf("failed to set status.selfLink: %v", err)
	}

	scheme := runtime.NewScheme()
	reader := fake.NewClientBuilder().WithScheme(scheme).WithObjects(readyKey).Build()

	t.Run("resolves KMS references by name", func(t *testing.T) {
		cluster := &krm.DataprocCluster{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-cluster",
				Namespace: "test-ns",
			},
			Spec: krm.DataprocClusterSpec{
				Config: &krm.ClusterConfig{
					EncryptionConfig: &krm.ClusterEncryptionConfig{
						GcePdKmsKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
							Name: "test-kms-key",
						},
					},
					SecurityConfig: &krm.ClusterSecurityConfig{
						KerberosConfig: &krm.ClusterKerberosConfig{
							KmsKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
								Name: "test-kms-key",
							},
						},
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, cluster, nil); err != nil {
			t.Fatalf("NormalizeReferences failed: %v", err)
		}

		if got := cluster.Spec.Config.EncryptionConfig.GcePdKmsKeyRef.External; got != selfLink {
			t.Errorf("GcePdKmsKeyRef.External = %q, want %q", got, selfLink)
		}
		if got := cluster.Spec.Config.SecurityConfig.KerberosConfig.KmsKeyRef.External; got != selfLink {
			t.Errorf("KmsKeyRef.External = %q, want %q", got, selfLink)
		}
	})

	t.Run("handles external KMS references directly", func(t *testing.T) {
		externalKey := "projects/other-project/locations/global/keyRings/other-keyring/cryptoKeys/external-key"
		cluster := &krm.DataprocCluster{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-cluster",
				Namespace: "test-ns",
			},
			Spec: krm.DataprocClusterSpec{
				Config: &krm.ClusterConfig{
					EncryptionConfig: &krm.ClusterEncryptionConfig{
						GcePdKmsKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
							External: externalKey,
						},
					},
					SecurityConfig: &krm.ClusterSecurityConfig{
						KerberosConfig: &krm.ClusterKerberosConfig{
							KmsKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
								External: externalKey,
							},
						},
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, cluster, nil); err != nil {
			t.Fatalf("NormalizeReferences failed: %v", err)
		}

		if got := cluster.Spec.Config.EncryptionConfig.GcePdKmsKeyRef.External; got != externalKey {
			t.Errorf("GcePdKmsKeyRef.External = %q, want %q", got, externalKey)
		}
		if got := cluster.Spec.Config.SecurityConfig.KerberosConfig.KmsKeyRef.External; got != externalKey {
			t.Errorf("KmsKeyRef.External = %q, want %q", got, externalKey)
		}
	})
}
