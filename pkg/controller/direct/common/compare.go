// Copyright 2024 Google LLC
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
	"fmt"
	"reflect"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Compare(x, y proto.Message) (*fieldmaskpb.FieldMask, error) {
	updateMask := &fieldmaskpb.FieldMask{}
	// 	return nil, fmt.Errorf("compare type not known, suggest write your own compare logic.")
	if x == nil || y == nil {
		return nil, fmt.Errorf("missing compare proto objects")
	}
	if reflect.TypeOf(x).Kind() == reflect.Ptr && x == y {
		return updateMask, nil
	}
	mx := x.ProtoReflect()
	my := y.ProtoReflect()
	if mx.IsValid() != my.IsValid() {
		return nil, fmt.Errorf("mismatch validity, at least one proto object is empty and read-only.")
	}
	vx := protoreflect.ValueOfMessage(mx)
	vy := protoreflect.ValueOfMessage(my)
	vx.Equal(vy)
	return updateMask, nil
}

func DeepEqual_StringAndTimestampPb(a string, b *timestamppb.Timestamp) bool {
	norm := NormalizeStringToTimestamp(a)
	if norm == nil {
		return b == nil
	}
	if b == nil {
		return false
	}
	x := norm.String()
	y := b.String()
	return x == y
}
