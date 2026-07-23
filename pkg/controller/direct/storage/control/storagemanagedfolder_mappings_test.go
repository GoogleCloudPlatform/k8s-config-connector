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

package storagecontrol

import (
	"testing"

	pb "cloud.google.com/go/storage/control/apiv2/controlpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// The KRM fuzzer cannot cover these: it declares ".name" unimplemented, so it
// never populates the field that guards the mapping below.

func TestStorageManagedFolderSpec_FromProto(t *testing.T) {
	mapCtx := &direct.MapContext{}
	in := &pb.ManagedFolder{Name: "projects/proj1/buckets/bucket1/managedfolders/mf1"}

	out := StorageManagedFolderSpec_FromProto(mapCtx, in)
	if mapCtx.Err() != nil {
		t.Fatalf("StorageManagedFolderSpec_FromProto returned error: %v", mapCtx.Err())
	}
	if out == nil {
		t.Fatal("StorageManagedFolderSpec_FromProto returned nil")
	}
	if out.StorageFolderParent == nil {
		t.Fatal("StorageFolderParent is nil; ProjectRef and StorageBucketRef promote through it")
	}
	if got, want := out.ProjectRef.External, "proj1"; got != want {
		t.Errorf("ProjectRef.External = %q, want %q", got, want)
	}
	// The bucket ref holds the bucket name, not the parent path.
	if got, want := out.StorageBucketRef.External, "bucket1"; got != want {
		t.Errorf("StorageBucketRef.External = %q, want %q", got, want)
	}
	if got, want := direct.ValueOf(out.ResourceID), "mf1"; got != want {
		t.Errorf("ResourceID = %q, want %q", got, want)
	}
}

func TestStorageManagedFolderSpec_FromProto_UnparseableName(t *testing.T) {
	mapCtx := &direct.MapContext{}
	// A name in the format the GCP API actually returns.
	in := &pb.ManagedFolder{Name: "projects/_/buckets/bucket1/managedFolders/mf1/"}

	out := StorageManagedFolderSpec_FromProto(mapCtx, in)
	if mapCtx.Err() == nil {
		t.Fatal("expected an error for an unparseable name")
	}
	if out != nil {
		t.Errorf("expected nil spec on parse error, got %+v", out)
	}
}
