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

package dataflow

import (
	"context"
	"testing"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataflow/v1beta1"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestDataflowJob_NormalizeReferences(t *testing.T) {
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

	t.Run("resolves KmsKeyRef by name", func(t *testing.T) {
		obj := &krm.DataflowJob{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-job",
				Namespace: "test-ns",
			},
			Spec: krm.DataflowJobSpec{
				KmsKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
					Name: "test-key",
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		if obj.Spec.KmsKeyRef.External != selfLink {
			t.Errorf("NormalizeReferences() External = %q, want %q", obj.Spec.KmsKeyRef.External, selfLink)
		}

		mapCtx := &direct.MapContext{}
		protoObj := DataflowJobSpec_ToProto(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("DataflowJobSpec_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj.Environment == nil || protoObj.Environment.ServiceKmsKeyName != selfLink {
			t.Errorf("protoObj Environment.ServiceKmsKeyName = %v, want %q", protoObj.Environment, selfLink)
		}

		fromProto := DataflowJobSpec_FromProto(mapCtx, protoObj)
		if mapCtx.Err() != nil {
			t.Fatalf("DataflowJobSpec_FromProto unexpected error: %v", mapCtx.Err())
		}
		if fromProto.KmsKeyRef == nil || fromProto.KmsKeyRef.External != selfLink {
			t.Errorf("fromProto KmsKeyRef.External = %v, want %q", fromProto.KmsKeyRef, selfLink)
		}
	})

	t.Run("handles external KmsKeyRef directly", func(t *testing.T) {
		externalKey := "projects/test-project/locations/us-central1/keyRings/test-keyring/cryptoKeys/external-key"
		obj := &krm.DataflowJob{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-job",
				Namespace: "test-ns",
			},
			Spec: krm.DataflowJobSpec{
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
		protoObj := DataflowJobSpec_ToProto(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("DataflowJobSpec_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj.Environment == nil || protoObj.Environment.ServiceKmsKeyName != externalKey {
			t.Errorf("protoObj Environment.ServiceKmsKeyName = %v, want %q", protoObj.Environment, externalKey)
		}
	})

	t.Run("fuzzer roundtrips successfully", func(t *testing.T) {
		f := dataflowJobFuzzer()
		for seed := int64(0); seed < 20; seed++ {
			f.FuzzSpec(t, seed)
		}
	})
}
