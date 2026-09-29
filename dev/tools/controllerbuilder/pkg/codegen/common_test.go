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

package codegen

import (
	"testing"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// TestMapsToGoStruct checks which proto messages are generated as a Go struct
// of their own. Timestamp is generated as a string and Struct as
// apiextensionsv1.JSON, so neither is a struct. ListValue has no special
// mapping, so it is generated as a struct. The walk in scaffold.ReferenceHints
// relies on this to avoid descending into a Timestamp and reporting fields like
// .seconds that the CRD does not have.
func TestMapsToGoStruct(t *testing.T) {
	for _, tc := range []struct {
		msg  protoreflect.MessageDescriptor
		want bool
	}{
		{(&timestamppb.Timestamp{}).ProtoReflect().Descriptor(), false},
		{(&structpb.Struct{}).ProtoReflect().Descriptor(), false},
		{(&structpb.ListValue{}).ProtoReflect().Descriptor(), true},
	} {
		t.Run(string(tc.msg.FullName()), func(t *testing.T) {
			// Act
			got := MapsToGoStruct(tc.msg)

			// Assert
			if got != tc.want {
				t.Errorf("MapsToGoStruct(%s) = %v, want %v", tc.msg.FullName(), got, tc.want)
			}
		})
	}
}
