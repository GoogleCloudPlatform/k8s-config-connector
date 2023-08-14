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

	"k8s.io/klog/v2"
)

type structField struct {
	f *reflect.StructField

	jsonKey string
}

var _ Field = &structField{}

func (f *structField) ID() FieldID {
	fieldID := ToFieldID(f.jsonKey)
	return fieldID
}

// JSONKey returns the key we will normally use for JSON serialization.
func (f *structField) JSONKey() string {
	return f.jsonKey
}

func (f *structField) Type() *reflectType {
	return allTypes.get(f.f.Type)
}

func (f *structField) isRequired() bool {
	// AWS uses required
	requiredTag := f.f.Tag.Get("required")
	switch requiredTag {
	case "true":
		return true

	case "false":
		return false

	case "":
		return false

	default:
		klog.Fatalf("unexpected required value %q", requiredTag)
	}

	return false
}

func (f *structField) getValue(p *point) *point {
	structVal := p.rv
	if structVal.Kind() == reflect.Ptr {
		structVal = structVal.Elem()
	}
	rv := structVal.FieldByIndex(f.f.Index)
	out := p.scope.newPoint(rv)
	return out
}

func (f *structField) setValue(p *point, srcVal reflect.Value) error {
	structVal := p.rv
	if structVal.Kind() == reflect.Ptr {
		structVal = structVal.Elem()
	}

	fieldVal := structVal.FieldByIndex(f.f.Index)
	destType := fieldVal.Type()
	destVal, err := p.scope.convert(srcVal, destType)
	if err != nil {
		return fmt.Errorf("converting %v to %v: %w", srcVal.Type(), destType, err)
	}
	if !destVal.IsValid() {
		return nil
	}
	fieldVal.Set(destVal)
	return nil
}
