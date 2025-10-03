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

package mockmonitoring

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func setField(dest proto.Message, src proto.Message, path string) error {
	tokens := strings.Split(path, ".")
	if len(tokens) == 0 {
		return fmt.Errorf("path is not valid: %q", path)
	}

	val, err := getFieldValue(src.ProtoReflect(), tokens)
	if err != nil {
		return err
	}

	if err := setFieldValue(dest.ProtoReflect(), tokens, val); err != nil {
		return err
	}
	return nil
}

func getFieldValue(src protoreflect.Message, path []string) (protoreflect.Value, error) {
	remainder := path
	for {
		token := remainder[0]
		remainder = remainder[1:]

		srcType := src.Descriptor()
		srcField := srcType.Fields().ByJSONName(token)
		if srcField == nil {
			return protoreflect.Value{}, fmt.Errorf("field %q not found in path %q", token, strings.Join(path, "."))
		}
		srcValue := src.Get(srcField)
		if len(remainder) == 0 {
			return srcValue, nil
		}
		switch srcValue := srcValue.Interface().(type) {
		case protoreflect.Message:
			src = srcValue

		default:
			return protoreflect.Value{}, fmt.Errorf("unhandled type %T", srcValue)
		}
	}
}

func setFieldValue(dest protoreflect.Message, path []string, val protoreflect.Value) error {
	remainder := path
	for {
		token := remainder[0]
		remainder = remainder[1:]

		destType := dest.Descriptor()
		destField := destType.Fields().ByJSONName(token)
		if destField == nil {
			return fmt.Errorf("field %q not found in path %q", token, strings.Join(path, "."))
		}
		if len(remainder) == 0 {
			dest.Set(destField, val)
			return nil
		}
		destValue := dest.Mutable(destField)
		switch destValue := destValue.Interface().(type) {
		case protoreflect.Message:
			dest = destValue

		default:
			return fmt.Errorf("unhandled type %T", destValue)
		}
	}
}
