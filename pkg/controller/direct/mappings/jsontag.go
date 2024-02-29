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
)

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

	// locationName is used by AWS types
	// TODO: Make this conditional?
	locationName := f.Tag.Get("locationName")
	if locationName != "" {
		return locationName
	}
	name := f.Name
	return goFieldToFieldID(name)
}
