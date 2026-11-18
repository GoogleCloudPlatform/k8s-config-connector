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

package sqladmin

import (
	"testing"

	pb "cloud.google.com/go/sql/apiv1/sqlpb"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sqladmin/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func TestDiskEncryptionConfigurationMapping(t *testing.T) {
	t.Run("nil input", func(t *testing.T) {
		mapCtx := &direct.MapContext{}
		if got := DiskEncryptionConfiguration_FromProto(mapCtx, nil); got != nil {
			t.Errorf("expected nil, got %+v", got)
		}
		if got := DiskEncryptionConfiguration_ToProto(mapCtx, nil); got != nil {
			t.Errorf("expected nil, got %+v", got)
		}
	})

	t.Run("empty input", func(t *testing.T) {
		mapCtx := &direct.MapContext{}
		fromProto := DiskEncryptionConfiguration_FromProto(mapCtx, &pb.DiskEncryptionConfiguration{})
		if fromProto == nil {
			t.Fatal("expected non-nil output")
		}
		if fromProto.KMSKeyRef != nil {
			t.Errorf("expected nil KMSKeyRef, got %+v", fromProto.KMSKeyRef)
		}

		toProto := DiskEncryptionConfiguration_ToProto(mapCtx, fromProto)
		if toProto == nil {
			t.Fatal("expected non-nil output")
		}
		if toProto.KmsKeyName != "" {
			t.Errorf("expected empty KmsKeyName, got %q", toProto.KmsKeyName)
		}
	})

	t.Run("roundtrip populated key", func(t *testing.T) {
		mapCtx := &direct.MapContext{}
		keyName := "projects/test-project/locations/us-central1/keyRings/my-ring/cryptoKeys/my-key"

		inProto := &pb.DiskEncryptionConfiguration{
			KmsKeyName: keyName,
		}
		krmObj := DiskEncryptionConfiguration_FromProto(mapCtx, inProto)
		if krmObj == nil || krmObj.KMSKeyRef == nil {
			t.Fatalf("expected non-nil KMSKeyRef, got %+v", krmObj)
		}
		if krmObj.KMSKeyRef.External != keyName {
			t.Errorf("expected External=%q, got %q", keyName, krmObj.KMSKeyRef.External)
		}

		outProto := DiskEncryptionConfiguration_ToProto(mapCtx, krmObj)
		if outProto == nil {
			t.Fatal("expected non-nil proto output")
		}
		if outProto.KmsKeyName != keyName {
			t.Errorf("expected KmsKeyName=%q, got %q", keyName, outProto.KmsKeyName)
		}
	})

	t.Run("krm to proto with nil or empty External", func(t *testing.T) {
		mapCtx := &direct.MapContext{}
		krmObj := &krm.DiskEncryptionConfiguration{
			KMSKeyRef: &kmsv1beta1.KMSCryptoKeyRef{},
		}
		outProto := DiskEncryptionConfiguration_ToProto(mapCtx, krmObj)
		if outProto == nil {
			t.Fatal("expected non-nil proto output")
		}
		if outProto.KmsKeyName != "" {
			t.Errorf("expected empty KmsKeyName, got %q", outProto.KmsKeyName)
		}
	})
}
