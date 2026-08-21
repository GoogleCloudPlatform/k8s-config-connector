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

package mappings

import (
	"reflect"
	"strings"
	"unicode"
)

// getJSONFieldTag is a helper to extract the json key for a field, from the reflection StructField.
func getJSONFieldTag(f *reflect.StructField) string {
	// Handle protobuf tags.
	// example: `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3"
	protoTag := f.Tag.Get("protobuf")
	for _, v := range strings.Split(protoTag, ",") {
		if strings.HasPrefix(v, "json=") {
			return strings.TrimPrefix(v, "json=")
		}
	}

	jsonTag := f.Tag.Get("json")
	switch jsonTag {
	case "-":
		// Not marshaled
		return ""

	case "":
		break

	default:
		parts := strings.Split(jsonTag, ",")
		name := parts[0]
		if name == "" {
			name = f.Name
		}
		return name
	}

	name := f.Name
	return fallbackGoFieldToFieldID(name)
}

// fallbackGoFieldToFieldID converts a go field into the json equivalent.
func fallbackGoFieldToFieldID(fieldName string) string {
	var out []rune
	for i, r := range fieldName {
		if i == 0 {
			r = unicode.ToLower(r)
		}
		out = append(out, r)
	}
	s := string(out)
	return s
}
