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

package common

import (
	"testing"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func TestBuildParentMap_Dynamic(t *testing.T) {
	// Let's create an actualPb using FileDescriptorProto which is deeply nested.
	// Options is of type *descriptorpb.FileOptions
	actualPb := &descriptorpb.FileDescriptorProto{
		Options: &descriptorpb.FileOptions{
			GoPackage: proto.String("my-package"),
		},
	}

	desiredPb := &descriptorpb.FileDescriptorProto{}

	paths := []string{
		"Options.GoPackage",
	}

	parentMap := BuildParentMap(desiredPb, actualPb, paths)

	// We expect "Options" to be created in desiredPb and registered in parentMap.
	pair, exists := parentMap["Options"]
	if !exists {
		t.Fatalf("expected 'Options' to be registered in parentMap")
	}

	if pair.Actual == nil {
		t.Errorf("expected Actual to be non-nil")
	}

	if pair.Desired == nil {
		t.Errorf("expected Desired to be non-nil")
	}

	// Verify that Options in desiredPb is not nil anymore (initialized)
	if desiredPb.Options == nil {
		t.Errorf("expected Options to be initialized on desiredPb")
	}
}

func TestFindProtoField(t *testing.T) {
	desc := (&descriptorpb.FileDescriptorProto{}).ProtoReflect().Descriptor()

	fd := FindProtoField(desc, "options")
	if fd == nil {
		t.Fatalf("expected to find field 'options'")
	}
	if fd.Name() != "options" {
		t.Errorf("expected field name 'options', got %q", fd.Name())
	}

	fdRef := FindProtoField(desc, "optionsRef")
	if fdRef == nil {
		t.Fatalf("expected to find field 'optionsRef' matching 'options'")
	}
	if fdRef.Name() != "options" {
		t.Errorf("expected field name 'options' for optionsRef, got %q", fdRef.Name())
	}
}
