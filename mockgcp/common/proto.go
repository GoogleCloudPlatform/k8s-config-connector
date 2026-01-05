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

package common

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"k8s.io/klog/v2"
)

func ProtoClone[T proto.Message](obj T) T {
	return proto.Clone(obj).(T)
}

// GetProtoValue retrieves the value of the given field path from the proto message.
// The path is a dot-separated string of json field names.
// It returns the value and a boolean indicating whether the field was found.
func GetProtoValue(msg proto.Message, path string) (protoreflect.Value, bool) {
	r := msg.ProtoReflect()
	fields := strings.Split(path, ".")
	for i, field := range fields {
		fd := r.Descriptor().Fields().ByJSONName(field)
		if fd == nil {
			return protoreflect.Value{}, false
		}
		v := r.Get(fd)
		if i == len(fields)-1 {
			return v, true
		}
		if fd.Message() == nil {
			return v, false
		}
		r = v.Message()
	}
	panic("should be unreachable")
}

// SetProtoValue sets the value of the given field path in the proto message.
// The path is a dot-separated string of json field names.
func SetProtoValue[T proto.Message](msg T, path string, value protoreflect.Value) error {
	r := msg.ProtoReflect()
	fields := strings.Split(path, ".")
	for i, field := range fields {
		fd := r.Descriptor().Fields().ByJSONName(field)
		if fd == nil {
			return fmt.Errorf("field %q not found in message %q", field, r.Descriptor().FullName())
		}
		if i == len(fields)-1 {
			r.Set(fd, value)
			return nil
		}
		v := r.Get(fd)
		if fd.Message() == nil {
			return fmt.Errorf("field %q in message %q is not a message", field, r.Descriptor().FullName())
		}
		r = v.Message()
	}
	panic("should be unreachable")
}

// MustCopyField copies the value of the given field from src to dst proto messages.
// It panics if the field is not found or cannot be set.
func MustCopyField(src proto.Message, dst proto.Message, field string) {
	v, ok := GetProtoValue(src, field)
	if !ok {
		klog.Fatalf("field %q not found in message %q", field, src.ProtoReflect().Descriptor().FullName())
	}
	if err := SetProtoValue(dst, field, v); err != nil {
		klog.Fatalf("setting field %q in message %q: %v", field, dst.ProtoReflect().Descriptor().FullName(), err)
	}
}
