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

package fields

import (
	"context"
	"sort"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// ComputeImpliedFieldMask computes the implied field mask when it is omitted.
// According to AIP 134: this is "equivalent to all fields that are populated (have a non-empty value)."
// https://google.aip.dev/134
func ComputeImpliedFieldMask(ctx context.Context, req proto.Message) []string {
	var fieldMask []string
	req.ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		fieldName := string(fd.Name())
		fieldMask = append(fieldMask, fieldName)
		return true
	})
	sort.Strings(fieldMask)
	return fieldMask
}
