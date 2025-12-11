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

package fields

import (
	"errors"
	"fmt"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"k8s.io/klog/v2"
)

// UpdateByFieldMask updates the `original` Message with the `update` Message value in the given `updatePaths` fields
func UpdateByFieldMask(original, update proto.Message, updatePaths []string) error {
	var errs []error
	for _, path := range updatePaths {
		if err := walk(original, update, path); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

func walk(original, update proto.Message, path string) error {
	originalRm := original.ProtoReflect()
	updateRm := update.ProtoReflect()
	segments := strings.SplitN(path, ".", 2)
	if len(segments) == 1 {
		return replace(originalRm, updateRm, segments[0])
	}
	return walk(originalChildMessage(originalRm, segments[0]), updateChildMessage(updateRm, segments[0]), segments[1])
}

func replace(original, update protoreflect.Message, fieldName string) error {
	originalFd := getFieldDescriptor(original.Descriptor().Fields(), fieldName)
	if originalFd == nil {
		return fmt.Errorf("field %q not found in %s", fieldName, original.Descriptor().FullName())
	}
	updateFd := getFieldDescriptor(update.Descriptor().Fields(), fieldName)
	updateVal := update.Get(updateFd)

	// Update Map
	if originalFd.IsMap() {
		originalVal := original.Mutable(originalFd)
		originalVal.Map().Range(func(k protoreflect.MapKey, v protoreflect.Value) bool {
			originalVal.Map().Clear(k)
			return true
		})
		updateVal.Map().Range(func(k protoreflect.MapKey, v protoreflect.Value) bool {
			originalVal.Map().Set(k, v)
			return true
		})
		return nil
	}
	// Update List
	if originalFd.IsList() {
		originalVal := original.Mutable(originalFd)
		originalVal.List().Truncate(0)
		for i := 0; i < updateVal.List().Len(); i++ {
			originalVal.List().Append(updateVal.List().Get(i))
		}
		return nil
	}

	switch originalFd.Kind() {
	case protoreflect.MessageKind, protoreflect.GroupKind:
		if !original.IsValid() {
			return fmt.Errorf("%s is read-only or empty", fieldName)
		}
		if update.Has(updateFd) {
			original.Set(updateFd, updateVal)
		} else {
			original.Clear(originalFd)
		}
		return nil
	case protoreflect.StringKind, protoreflect.DoubleKind, protoreflect.Int32Kind, protoreflect.Int64Kind, protoreflect.Uint64Kind, protoreflect.BoolKind, protoreflect.EnumKind:
		if !original.IsValid() {
			return fmt.Errorf("%s is read-only or empty", fieldName)
		}
		original.Set(updateFd, updateVal)
		return nil
	default:
		klog.Warningf("unhandled type %v for field %v", originalFd.Kind(), fieldName)
		return fmt.Errorf("unhandled type %v for field %v", originalFd.Kind(), fieldName)
	}
}

// originalChildMessage get the orignal Message's mutable reference to the `fieldName“ composite.
func originalChildMessage(m protoreflect.Message, fieldName string) proto.Message {
	fd := getFieldDescriptor(m.Descriptor().Fields(), fieldName)
	if fd == nil {
		// Panic or return nil? The caller expects a message.
		// If we return nil, next call might panic.
		// We'll panic with a clear message.
		panic(fmt.Errorf("field %q not found in %s", fieldName, m.Descriptor().FullName()))
	}
	return m.Mutable(fd).Message().Interface()
}

func updateChildMessage(m protoreflect.Message, fieldName string) proto.Message {
	fd := getFieldDescriptor(m.Descriptor().Fields(), fieldName)
	if fd == nil {
		panic(fmt.Errorf("field %q not found in %s", fieldName, m.Descriptor().FullName()))
	}
	return m.Get(fd).Message().Interface()
}

func getFieldDescriptor(fields protoreflect.FieldDescriptors, name string) protoreflect.FieldDescriptor {
	if fd := fields.ByName(protoreflect.Name(name)); fd != nil {
		return fd
	}
	return fields.ByJSONName(name)
}
