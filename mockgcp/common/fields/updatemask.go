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
)

func UpdateByFieldMask(old, new proto.Message, updatePaths []string) error {
	var errs []error
	for _, path := range updatePaths {
		if err := walk(old, new, path); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

func walk(old, new proto.Message, path string) error {
	segments := strings.Split(path, ".")
	if len(segments) == 1 {
		return replace(old, new, segments[0])
	}

	subPath := strings.TrimPrefix(path, segments[0]+".")
	return walk(childMessage(old, segments[0]), childMessage(new, segments[0]), subPath)
}

func replace(old, new proto.Message, fieldName string) error {
	oldFd := old.ProtoReflect().Descriptor().Fields().ByJSONName(fieldName)
	oldVal := old.ProtoReflect().Get(oldFd)
	newFd := new.ProtoReflect().Descriptor().Fields().ByJSONName(fieldName)
	newVal := new.ProtoReflect().Get(newFd)

	if oldFd.Kind() != protoreflect.MessageKind {
		m := old.ProtoReflect()
		if !m.IsValid() {
			return fmt.Errorf("%s is read-only or empty", fieldName)
		}
		m.Set(newFd, newVal)
	}
	// Update Map
	if oldFd.IsMap() {
		newVal.Map().Range(func(k protoreflect.MapKey, v protoreflect.Value) bool {
			oldVal.Map().Set(k, v)
			return true
		})
	}
	// Update List
	if oldFd.IsList() {
		oldVal.List().Truncate(0)
		for i := 0; i < newVal.List().Len(); i++ {
			oldVal.List().Append(newVal.List().Get(i))
		}
	}
	return nil
}

func childMessage(m proto.Message, fieldName string) proto.Message {
	fd := m.ProtoReflect().Descriptor().Fields().ByJSONName(fieldName)
	v := m.ProtoReflect().Get(fd)
	return v.Message().Interface()
}
