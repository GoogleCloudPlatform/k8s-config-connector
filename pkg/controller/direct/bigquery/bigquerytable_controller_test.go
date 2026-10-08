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

package bigquery

import (
	"context"
	"testing"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1beta1"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"google.golang.org/api/googleapi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestAdapter_IsUnreadableButDeletable(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "nil error",
			err:  nil,
			want: false,
		},
		{
			name: "HTTP 400 Bad Request api error",
			err: &googleapi.Error{
				Code:    400,
				Message: "Bad Request",
			},
			want: true,
		},
		{
			name: "HTTP 404 Not Found api error",
			err: &googleapi.Error{
				Code:    404,
				Message: "Not Found",
			},
			want: false,
		},
		{
			name: "HTTP 403 Forbidden api error",
			err: &googleapi.Error{
				Code:    403,
				Message: "Forbidden",
			},
			want: false,
		},
		{
			name: "HTTP 500 Internal Server api error",
			err: &googleapi.Error{
				Code:    500,
				Message: "Internal Server Error",
			},
			want: false,
		},
		{
			name: "grpc InvalidArgument",
			err:  status.Error(codes.InvalidArgument, "invalid argument"),
			want: false,
		},
	}

	a := &Adapter{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := a.IsUnreadableButDeletable(tt.err)
			if got != tt.want {
				t.Errorf("IsUnreadableButDeletable() = %v, want %v for error: %v", got, tt.want, tt.err)
			}
		})
	}
}

func TestBigQueryTable_NormalizeReferences(t *testing.T) {
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

	t.Run("resolves KmsKeyRef by name", func(t *testing.T) {
		obj := &krm.BigQueryTable{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-table",
				Namespace: "test-ns",
			},
			Spec: krm.BigQueryTableSpec{
				EncryptionConfiguration: &krm.TableEncryptionConfiguration{
					KmsKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
						Name: "test-key",
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		if obj.Spec.EncryptionConfiguration.KmsKeyRef.External != selfLink {
			t.Errorf("NormalizeReferences() External = %q, want %q", obj.Spec.EncryptionConfiguration.KmsKeyRef.External, selfLink)
		}

		mapCtx := &direct.MapContext{}
		protoObj := EncryptionConfiguration_ToProto(mapCtx, obj.Spec.EncryptionConfiguration)
		if mapCtx.Err() != nil {
			t.Fatalf("EncryptionConfiguration_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj.KmsKeyName != selfLink {
			t.Errorf("protoObj KmsKeyName = %q, want %q", protoObj.KmsKeyName, selfLink)
		}

		fromProto := EncryptionConfiguration_FromProto(mapCtx, protoObj)
		if mapCtx.Err() != nil {
			t.Fatalf("EncryptionConfiguration_FromProto unexpected error: %v", mapCtx.Err())
		}
		if fromProto.KmsKeyRef == nil || fromProto.KmsKeyRef.External != selfLink {
			t.Errorf("fromProto KmsKeyRef.External = %v, want %q", fromProto.KmsKeyRef, selfLink)
		}
	})

	t.Run("handles external KmsKeyRef directly", func(t *testing.T) {
		externalKey := "projects/test-project/locations/us-central1/keyRings/test-keyring/cryptoKeys/external-key"
		obj := &krm.BigQueryTable{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-table",
				Namespace: "test-ns",
			},
			Spec: krm.BigQueryTableSpec{
				EncryptionConfiguration: &krm.TableEncryptionConfiguration{
					KmsKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
						External: externalKey,
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		if obj.Spec.EncryptionConfiguration.KmsKeyRef.External != externalKey {
			t.Errorf("NormalizeReferences() External = %q, want %q", obj.Spec.EncryptionConfiguration.KmsKeyRef.External, externalKey)
		}

		mapCtx := &direct.MapContext{}
		protoObj := EncryptionConfiguration_ToProto(mapCtx, obj.Spec.EncryptionConfiguration)
		if mapCtx.Err() != nil {
			t.Fatalf("EncryptionConfiguration_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj.KmsKeyName != externalKey {
			t.Errorf("protoObj KmsKeyName = %q, want %q", protoObj.KmsKeyName, externalKey)
		}

		fromProto := EncryptionConfiguration_FromProto(mapCtx, protoObj)
		if mapCtx.Err() != nil {
			t.Fatalf("EncryptionConfiguration_FromProto unexpected error: %v", mapCtx.Err())
		}
		if fromProto.KmsKeyRef == nil || fromProto.KmsKeyRef.External != externalKey {
			t.Errorf("fromProto KmsKeyRef.External = %v, want %q", fromProto.KmsKeyRef, externalKey)
		}
	})

	t.Run("handles nil EncryptionConfiguration", func(t *testing.T) {
		obj := &krm.BigQueryTable{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-table",
				Namespace: "test-ns",
			},
			Spec: krm.BigQueryTableSpec{},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		mapCtx := &direct.MapContext{}
		protoObj := EncryptionConfiguration_ToProto(mapCtx, obj.Spec.EncryptionConfiguration)
		if mapCtx.Err() != nil {
			t.Fatalf("EncryptionConfiguration_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj != nil {
			t.Errorf("protoObj = %v, want nil", protoObj)
		}
	})
}
