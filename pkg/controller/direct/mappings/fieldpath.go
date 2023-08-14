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
	"fmt"
	"reflect"
	"strings"
)

// fieldPath is a sequence of field ids, forming a nested path such as `spec.size`
type fieldPath struct {
	fields []FieldID
}

// newFieldPath builds a fieldPath.
func newFieldPath(fields ...FieldID) *fieldPath {
	return &fieldPath{fields: fields}
}

// String returns a user-friendly value.
func (f *fieldPath) String() string {
	var s []string
	for _, v := range f.fields {
		s = append(s, string(v))
	}
	return strings.Join(s, ".")
}

// FindPoint will return a point representing a nested child field of parent
func (f *fieldPath) FindPoint(parent *point) *point {
	p := parent
	for _, field := range f.fields {
		p = p.Child(field)
	}
	return p
}

// SetValue will set the nested field in parent to the value v
func (f *fieldPath) SetValue(parent *point, v reflect.Value) error {
	p := parent
	n := len(f.fields)
	for i := 0; i < n-1; i++ {
		p = p.Child(f.fields[i])
		if p == nil {
			return fmt.Errorf("unable to find path %v", f)
		}
	}
	return p.SetValue(f.fields[n-1], v)
}

// parseFieldPath builds a new fieldPath
func parseFieldPath(s string) *fieldPath {
	var out []FieldID
	for _, v := range strings.Split(s, ".") {
		out = append(out, toFieldID(v))
	}
	return newFieldPath(out...)
}
