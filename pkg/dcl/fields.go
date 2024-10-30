// Copyright 2022 Google LLC
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

package dcl

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/pathslice"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/slice"
)

func IsContainerField(path []string) bool {
	if len(path) > 1 {
		return false
	}
	field := pathslice.Base(path)
	return field == "organization" || field == "project" || field == "folder"
}

// TrimNilFields removes all nil fields in the input object, including
// non-top-level nil fields (e.g. those found in nested objects and nested list
// of objects).
func TrimNilFields(m map[string]interface{}) {
	for k, v := range m {
		if IsNil(v) {
			delete(m, k)
			continue
		}

		switch v := v.(type) {
		// If value is an object, trim nil fields in object.
		case map[string]interface{}:
			TrimNilFields(v)
		// If value is a list of objects, trim nil fields in each object.
		case []interface{}:
			if !slice.IsListOfStringInterfaceMaps(v) {
				// List is a list of primitives rather than of objects.
				continue
			}
			for _, e := range v {
				TrimNilFields(e.(map[string]interface{}))
			}
		}
	}
}

// An interface value is strictly equal to nil only if the V (actual value) and T (type) are both unset;
// we want to treat interface values like (T=[]interface{}, V=nil) as nil too since the actual value is nil.
func IsNil(v interface{}) bool {
	if v == nil {
		return true
	}
	switch v := v.(type) {
	case []interface{}:
		return v == nil
	case map[string]interface{}:
		return v == nil
	}
	return false
}

func AddToMap(key string, val interface{}, obj map[string]interface{}) {
	if !IsNil(val) {
		obj[key] = val
	}
}
