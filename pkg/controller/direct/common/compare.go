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
	"k8s.io/apimachinery/pkg/util/sets"
)

type CompareDiff func(fieldName protoreflect.Name, a, b proto.Message) (bool, error)

var BasicDiff = func(fieldName protoreflect.Name, a, b proto.Message) (bool, error) {
	aField := a.ProtoReflect().Descriptor().Fields().ByName(fieldName)
	bField := b.ProtoReflect().Descriptor().Fields().ByName(fieldName)

	// Skip output-only fields
	if IsFieldBehavior(aField, annotations.FieldBehavior_OUTPUT_ONLY) {
		return false, nil
	}
	// The field is previously unset
	if bField == nil {
		return true, nil
	}
	if !aField.Kind().IsValid() {
		return false, fmt.Errorf("unimplemented kind: %s", aField.Kind().String())
	}

	diff := false
	aVal := a.ProtoReflect().Get(aField)
	bVal := b.ProtoReflect().Get(bField)
	switch aField.Kind() {
	case protoreflect.MessageKind:
		if aField.IsList() || aField.IsMap() {
			if !aVal.Equal(bVal) {
				diff = true
			}
		} else {
			m := aVal.Message().Interface()
			// Compare well-known proto type as a whole otherwise the diffPath (update field mask) could be wrong.
			switch m.(type) {
			case *timestamppb.Timestamp:
				if !aVal.Equal(bVal) {
					diff = true
				}
			case *durationpb.Duration:
				if !aVal.Equal(bVal) {
					diff = true
				}
			default:
				return false, fmt.Errorf("field %s not recursed", fieldName)
			}
		}
	default:
		if !aVal.Equal(bVal) {
			diff = true
		}
	}
	if diff && IsFieldBehavior(aField, annotations.FieldBehavior_IMMUTABLE) {
		return false, fmt.Errorf("change to immutable field %q", fieldName)
	}
	return diff, nil
}

func CompareProtoMessage(a, b proto.Message, compareDiff CompareDiff) (sets.Set[string], error) {
	diffPaths := sets.Set[string]{}
	aDescriptor := a.ProtoReflect().Descriptor()

	for i := 0; i < aDescriptor.Fields().Len(); i++ {
		field := aDescriptor.Fields().Get(i)
		updatePath := updatePathFromField(a, aDescriptor.Fields().Get(i))

		aVal := a.ProtoReflect().Get(field)
		bVal := b.ProtoReflect().Get(field)
		if shouldRecurse(field, a.ProtoReflect()) {
			subPaths, err := CompareProtoMessage(aVal.Message().Interface(), bVal.Message().Interface(), compareDiff)
			if err != nil {
				return nil, err
			}
			for path := range subPaths {
				diffPaths.Insert(updatePath + "." + path)
			}
		} else {
			if diff, err := compareDiff(field.Name(), a, b); err != nil {
				return nil, err
			} else if diff {
				diffPaths.Insert(updatePath)
			}
		}
	}
	return diffPaths, nil
}

func shouldRecurse(field protoreflect.FieldDescriptor, message protoreflect.Message) bool {
	if field.Kind() != protoreflect.MessageKind {
		return false
	}
	if field.IsList() || field.IsMap() {
		return false
	}
	m := message.Get(field).Message().Interface()
	// Compare well-known proto type as a whole otherwise the diffPath (update field mask) could be wrong.
	switch m.(type) {
	case *timestamppb.Timestamp:
		return false
	case *durationpb.Duration:
		return false
	default:
		return true
	}
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
