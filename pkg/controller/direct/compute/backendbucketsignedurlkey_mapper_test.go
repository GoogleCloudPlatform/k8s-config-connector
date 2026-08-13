// Copyright 2025 Google LLC
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

package compute

import (
	"context"
	"testing"

	computepb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	refsecret "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
	"k8s.io/apimachinery/pkg/runtime"
)

func TestSignedURLKey_ToProto(t *testing.T) {
	proto := SignedURLKey_ToProto("my-key", "abc123==")
	if proto == nil {
		t.Fatal("expected non-nil SignedUrlKey")
	}
	if proto.GetKeyName() != "my-key" {
		t.Errorf("KeyName: got %q, want my-key", proto.GetKeyName())
	}
	if proto.GetKeyValue() != "abc123==" {
		t.Errorf("KeyValue: got %q, want abc123==", proto.GetKeyValue())
	}
}

func TestSignedURLKey_KeyNameFromProto_NonNil(t *testing.T) {
	name := "test-key"
	proto := &computepb.SignedUrlKey{KeyName: &name}
	got := SignedURLKey_KeyNameFromProto(proto)
	if got != "test-key" {
		t.Errorf("got %q, want test-key", got)
	}
}

func TestSignedURLKey_KeyNameFromProto_Nil(t *testing.T) {
	got := SignedURLKey_KeyNameFromProto(nil)
	if got != "" {
		t.Errorf("got %q, want empty string", got)
	}
}

func TestParseBackendBucketSignedURLKeyExternal_Valid(t *testing.T) {
	external := "projects/my-project/global/backendBuckets/my-bucket/signedUrlKeys/my-key"
	id, err := krm.ParseBackendBucketSignedURLKeyExternal(external)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if id.String() != external {
		t.Errorf("String(): got %q, want %q", id.String(), external)
	}
	if id.Parent().ProjectID != "my-project" {
		t.Errorf("ProjectID: got %q, want my-project", id.Parent().ProjectID)
	}
	if id.Parent().BackendBucket != "my-bucket" {
		t.Errorf("BackendBucket: got %q, want my-bucket", id.Parent().BackendBucket)
	}
	if id.KeyName() != "my-key" {
		t.Errorf("KeyName: got %q, want my-key", id.KeyName())
	}
}

func TestParseBackendBucketSignedURLKeyExternal_Errors(t *testing.T) {
	cases := []struct {
		name  string
		input string
	}{
		{"too few tokens", "projects/p/global/backendBuckets/bucket/signedUrlKeys"},
		{"wrong structure", "notprojects/p/global/backendBuckets/bucket/signedUrlKeys/key"},
		{"wrong key segment", "projects/p/global/backendBuckets/bucket/wrongSegment/key"},
		{"empty project", "projects//global/backendBuckets/bucket/signedUrlKeys/key"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := krm.ParseBackendBucketSignedURLKeyExternal(tc.input)
			if err == nil {
				t.Errorf("ParseBackendBucketSignedURLKeyExternal(%q): expected error, got nil", tc.input)
			}
		})
	}
}

func TestBackendBucketSignedURLKeyIdentity_String(t *testing.T) {
	external := "projects/p/global/backendBuckets/my-bucket/signedUrlKeys/key1"
	id, err := krm.ParseBackendBucketSignedURLKeyExternal(external)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if id.String() != external {
		t.Errorf("got %q, want %q", id.String(), external)
	}
}

func TestExport_BackendBucketSignedURLKey(t *testing.T) {
	external := "projects/my-project/global/backendBuckets/my-bucket/signedUrlKeys/my-key"
	id, err := krm.ParseBackendBucketSignedURLKeyExternal(external)
	if err != nil {
		t.Fatalf("ParseBackendBucketSignedURLKeyExternal: %v", err)
	}
	a := &BackendBucketSignedURLKeyAdapter{
		id:    id,
		found: true,
	}
	u, err := a.Export(context.Background())
	if err != nil {
		t.Fatalf("Export() error: %v", err)
	}
	obj := &krm.ComputeBackendBucketSignedURLKey{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, obj); err != nil {
		t.Fatalf("FromUnstructured: %v", err)
	}
	if obj.Spec.ProjectRef == nil || obj.Spec.ProjectRef.External != "my-project" {
		t.Errorf("ProjectRef: got %v, want external=my-project", obj.Spec.ProjectRef)
	}
	if obj.Spec.BackendBucketRef.External != "my-bucket" {
		t.Errorf("BackendBucketRef: got %q, want my-bucket", obj.Spec.BackendBucketRef.External)
	}
	if obj.Spec.ResourceID == nil || *obj.Spec.ResourceID != "my-key" {
		t.Errorf("ResourceID: got %v, want my-key", obj.Spec.ResourceID)
	}
	// KeyValue should be empty (write-only in GCP)
	if obj.Spec.KeyValue.Value != nil {
		t.Errorf("KeyValue.Value should be nil for exported key (write-only), got %q", *obj.Spec.KeyValue.Value)
	}
}

func TestExport_BackendBucketSignedURLKey_NotFound(t *testing.T) {
	external := "projects/p/global/backendBuckets/b/signedUrlKeys/k"
	id, err := krm.ParseBackendBucketSignedURLKeyExternal(external)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	a := &BackendBucketSignedURLKeyAdapter{
		id:    id,
		found: false,
	}
	_, err = a.Export(context.Background())
	if err == nil {
		t.Error("Export() with found=false: expected error, got nil")
	}
}

func TestFuzzer_BackendBucketSignedURLKey_RoundTrip(t *testing.T) {
	keyName := "round-trip-key"
	keyValue := "dGVzdHZhbHVl"
	spec := &krm.ComputeBackendBucketSignedURLKeySpec{
		ProjectRef:       &refv1beta1.ProjectRef{External: "my-project"},
		BackendBucketRef: krm.BackendBucketRef{External: "my-bucket"},
		ResourceID:       &keyName,
		KeyValue:         refsecret.Legacy{Value: &keyValue},
	}

	// spec → proto
	proto := computeBackendBucketSignedURLKeySpec_ToProto(nil, spec)
	if proto.GetKeyName() != keyName {
		t.Errorf("ToProto KeyName: got %q, want %q", proto.GetKeyName(), keyName)
	}
	if proto.GetKeyValue() != keyValue {
		t.Errorf("ToProto KeyValue: got %q, want %q", proto.GetKeyValue(), keyValue)
	}

	// proto → spec
	got := computeBackendBucketSignedURLKeySpec_FromProto(nil, proto)
	if got.ResourceID == nil || *got.ResourceID != keyName {
		t.Errorf("FromProto ResourceID: got %v, want %q", got.ResourceID, keyName)
	}
	if got.KeyValue.Value == nil || *got.KeyValue.Value != keyValue {
		t.Errorf("FromProto KeyValue: got %v, want %q", got.KeyValue.Value, keyValue)
	}
}
