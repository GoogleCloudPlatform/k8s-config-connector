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

type fieldPath struct {
	fields []FieldID
}

func newFieldPath(fields ...FieldID) *fieldPath {
	return &fieldPath{fields: fields}
}

func (f *fieldPath) String() string {
	var s []string
	for _, v := range f.fields {
		s = append(s, string(v))
	}
	return strings.Join(s, ".")
}

func (f *fieldPath) FindPoint(p *point) *point {
	for _, field := range f.fields {
		p = p.Child(field)
	}
	return p
}

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

func ParseFieldPath(s string) *fieldPath {
	var out []FieldID
	for _, v := range strings.Split(s, ".") {
		out = append(out, ToFieldID(v))
	}
	return newFieldPath(out...)
}
