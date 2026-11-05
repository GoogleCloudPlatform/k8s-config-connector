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

package apihub

import (
	"context"
	"testing"

	pb "cloud.google.com/go/apihub/apiv1/apihubpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apihub/v1alpha1"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestAPIHubInstance_NormalizeReferences(t *testing.T) {
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

	t.Run("resolves KMSCryptoKeyRef by name", func(t *testing.T) {
		obj := &krm.APIHubInstance{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-instance",
				Namespace: "test-ns",
			},
			Spec: krm.APIHubInstanceSpec{
				Config: &krm.APIHubInstance_Config{
					CmekKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
						Name: "test-key",
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		if obj.Spec.Config.CmekKeyRef.External != selfLink {
			t.Errorf("NormalizeReferences() External = %q, want %q", obj.Spec.Config.CmekKeyRef.External, selfLink)
		}

		mapCtx := &direct.MapContext{}
		protoObj := APIHubInstance_Config_ToProto(mapCtx, obj.Spec.Config)
		if mapCtx.Err() != nil {
			t.Fatalf("APIHubInstance_Config_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj.GetCmekKeyName() != selfLink {
			t.Errorf("protoObj CmekKeyName = %q, want %q", protoObj.GetCmekKeyName(), selfLink)
		}

		fromProto := APIHubInstance_Config_FromProto(mapCtx, protoObj)
		if mapCtx.Err() != nil {
			t.Fatalf("APIHubInstance_Config_FromProto unexpected error: %v", mapCtx.Err())
		}
		if fromProto.CmekKeyRef == nil || fromProto.CmekKeyRef.External != selfLink {
			t.Errorf("fromProto CmekKeyRef.External = %v, want %q", fromProto.CmekKeyRef, selfLink)
		}
	})

	t.Run("handles external KMSCryptoKeyRef directly", func(t *testing.T) {
		externalKey := "projects/test-project/locations/us-central1/keyRings/test-keyring/cryptoKeys/external-key"
		obj := &krm.APIHubInstance{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-instance",
				Namespace: "test-ns",
			},
			Spec: krm.APIHubInstanceSpec{
				Config: &krm.APIHubInstance_Config{
					CmekKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
						External: externalKey,
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		if obj.Spec.Config.CmekKeyRef.External != externalKey {
			t.Errorf("NormalizeReferences() External = %q, want %q", obj.Spec.Config.CmekKeyRef.External, externalKey)
		}

		mapCtx := &direct.MapContext{}
		protoObj := APIHubInstance_Config_ToProto(mapCtx, obj.Spec.Config)
		if mapCtx.Err() != nil {
			t.Fatalf("APIHubInstance_Config_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj.GetCmekKeyName() != externalKey {
			t.Errorf("protoObj CmekKeyName = %q, want %q", protoObj.GetCmekKeyName(), externalKey)
		}
	})

	t.Run("returns error when referenced key is not found", func(t *testing.T) {
		obj := &krm.APIHubInstance{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-instance",
				Namespace: "test-ns",
			},
			Spec: krm.APIHubInstanceSpec{
				Config: &krm.APIHubInstance_Config{
					CmekKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
						Name: "nonexistent-key",
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err == nil {
			t.Fatalf("NormalizeReferences() expected error for nonexistent key, got nil")
		}
	})

	t.Run("handles nil CmekKeyRef gracefully", func(t *testing.T) {
		obj := &krm.APIHubInstance{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-instance",
				Namespace: "test-ns",
			},
			Spec: krm.APIHubInstanceSpec{
				Config: &krm.APIHubInstance_Config{},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		mapCtx := &direct.MapContext{}
		protoObj := APIHubInstance_Config_ToProto(mapCtx, obj.Spec.Config)
		if mapCtx.Err() != nil {
			t.Fatalf("APIHubInstance_Config_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj.GetCmekKeyName() != "" {
			t.Errorf("protoObj CmekKeyName = %q, want empty", protoObj.GetCmekKeyName())
		}

		fromProto := APIHubInstance_Config_FromProto(mapCtx, &pb.ApiHubInstance_Config{})
		if mapCtx.Err() != nil {
			t.Fatalf("APIHubInstance_Config_FromProto unexpected error: %v", mapCtx.Err())
		}
		if fromProto.CmekKeyRef != nil {
			t.Errorf("fromProto CmekKeyRef = %v, want nil", fromProto.CmekKeyRef)
		}
	})
}
