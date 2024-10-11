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
	"strings"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func CompareProtoMessage(a, b proto.Message, skipFields ...string) ([]string, error) {
	diffPaths := []string{}
	aDescriptor := a.ProtoReflect().Descriptor()
	bDescriptor := b.ProtoReflect().Descriptor()

	for i := 0; i < aDescriptor.Fields().Len(); i++ {
		aField := aDescriptor.Fields().Get(i)
		bField := bDescriptor.Fields().ByName(aField.Name())
		updatePath := updatePathFromField(a, aField)

		// Skip custom configured fields
		for _, skip := range skipFields {
			if skip == updatePath {
				continue
			}
		}
		// Skip output-only fields
		if IsFieldBehavior(aField, annotations.FieldBehavior_OUTPUT_ONLY) {
			continue
		}
		// The field is previously unset
		if bField == nil {
			diffPaths = append(diffPaths, updatePath)
			continue
		}
		if !aField.Kind().IsValid() {
			return nil, fmt.Errorf("unimplemented kind: " + aField.Kind().String())
		}

		aVal := a.ProtoReflect().Get(aField)
		bVal := b.ProtoReflect().Get(bField)
		var diffFPath []string

		switch aField.Kind() {
		case protoreflect.MessageKind:
			if aField.IsList() || aField.IsMap() {
				if !aVal.Equal(bVal) {
					diffFPath = append(diffFPath, updatePath)
				}
			} else {
				m := aVal.Message().Interface()
				// Compare well-known proto type as a whole otherwise the diffPath (update field mask) could be wrong.
				switch m.(type) {
				case *timestamppb.Timestamp:
					if !aVal.Equal(bVal) {
						diffFPath = append(diffFPath, updatePath)
					}
				case *durationpb.Duration:
					if !aVal.Equal(bVal) {
						diffFPath = append(diffFPath, updatePath)
					}
				default:
					var subSkipFields []string
					for _, skip := range skipFields {
						subSkipFields = append(subSkipFields, strings.TrimPrefix(skip, string(aField.Name())+"."))
					}
					subPaths, err := CompareProtoMessage(aVal.Message().Interface(), bVal.Message().Interface(), skipFields...)
					if err != nil {
						return nil, err
					}
					for _, path := range subPaths {
						diffFPath = append(diffFPath, updatePath+"."+path)
					}
				}
			}
		default:
			if !aVal.Equal(bVal) {
				diffFPath = append(diffFPath, updatePath)
			}
		}

		if len(diffFPath) != 0 {
			if IsFieldBehavior(aField, annotations.FieldBehavior_IMMUTABLE) {
				return nil, fmt.Errorf("change to immutable field %s", diffFPath)
			}
			diffPaths = append(diffPaths, diffFPath...)
		}
	}
	return diffPaths, nil
}

func updatePathFromField(obj proto.Message, field protoreflect.FieldDescriptor) string {
	d := obj.ProtoReflect().Descriptor()
	return strings.TrimPrefix(string(field.FullName()), string(d.FullName())+".")
}

func IsFieldBehavior(field protoreflect.FieldDescriptor, fieldBehavior annotations.FieldBehavior) bool {
	d := field.Options()
	fieldBehaviors := proto.GetExtension(d, annotations.E_FieldBehavior).([]annotations.FieldBehavior)
	for _, f := range fieldBehaviors {
		if f == fieldBehavior {
			return true
		}
	}
	return false
}
