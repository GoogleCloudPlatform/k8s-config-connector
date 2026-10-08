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

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataproc/v1alpha1"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestDataprocBatch_NormalizeReferences(t *testing.T) {
	ctx := context.Background()

	readyKey := &unstructured.Unstructured{}
	readyKey.SetGroupVersionKind(kmsv1beta1.KMSCryptoKeyGVK)
	readyKey.SetName("test-kms-key")
	readyKey.SetNamespace("test-ns")
	if err := unstructured.SetNestedField(readyKey.Object, "projects/test-project/locations/us-central1/keyRings/test-keyring/cryptoKeys/test-kms-key", "status", "selfLink"); err != nil {
		t.Fatalf("failed to set status.selfLink: %v", err)
	}

	scheme := runtime.NewScheme()
	reader := fake.NewClientBuilder().WithScheme(scheme).WithObjects(readyKey).Build()

	batch := &krm.DataprocBatch{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-batch",
			Namespace: "test-ns",
		},
		Spec: krm.DataprocBatchSpec{
			EnvironmentConfig: &krm.EnvironmentConfig{
				ExecutionConfig: &krm.ExecutionConfig{
					KMSKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
						Name: "test-kms-key",
					},
				},
			},
		},
	}

	if err := common.NormalizeReferences(ctx, reader, batch, nil); err != nil {
		t.Fatalf("NormalizeReferences failed: %v", err)
	}

	wantExternal := "projects/test-project/locations/us-central1/keyRings/test-keyring/cryptoKeys/test-kms-key"
	if got := batch.Spec.EnvironmentConfig.ExecutionConfig.KMSKeyRef.External; got != wantExternal {
		t.Errorf("KMSKeyRef.External = %q, want %q", got, wantExternal)
	}
}
