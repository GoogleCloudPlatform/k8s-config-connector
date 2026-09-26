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

// TestMapsToGoStruct pins which messages a caller walking the generated types
// may descend into. A Timestamp is written as a string, so a walk that entered
// it would report fields no generated struct has.
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
