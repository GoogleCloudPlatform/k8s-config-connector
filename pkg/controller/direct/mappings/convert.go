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
)

// convert will convert from the value src to destType, supporting assignment to a field of destType.
func (m *Mapping) convert(src reflect.Value, destType reflect.Type) (reflect.Value, error) {
	if src.Kind() == reflect.Pointer {
		if src.IsNil() {
			// Nothing to set
			return reflect.Value{}, nil
		}
		src = src.Elem()
	}

	if src.Kind() == reflect.Slice {
		if src.IsNil() {
			// Nothing to set
			return reflect.Value{}, nil
		}
		dest := reflect.New(destType).Elem()
		n := src.Len()
		for i := 0; i < n; i++ {
			srcElem := src.Index(i)
			destElem, err := m.convert(srcElem, destType.Elem())
			if err != nil {
				return reflect.Value{}, fmt.Errorf("converting slice element: %w", err)
			}
			// TODO: What if destElem is not valid
			dest = reflect.Append(dest, destElem)
		}
		return dest, nil
	}

	srcType := src.Type()

	switch srcType.String() {
	case "string":
		v := src.String()
		switch destType.String() {
		case "string":
			return reflect.ValueOf(v), nil
		case "*string":
			// When copying to an optional string, skip empty values
			if v == "" {
				return reflect.Value{}, nil
			}
			return reflect.ValueOf(&v), nil
		}
	}

	if src.CanInterface() {
		srcVal := src
		if srcVal.Kind() == reflect.Struct {
			if !srcVal.CanAddr() {
				return reflect.Value{}, fmt.Errorf("cannot address struct")
			}
			srcVal = srcVal.Addr()
		}
		var destVal reflect.Value
		if destType.Kind() == reflect.Pointer {
			destVal = reflect.New(destType.Elem())
		} else {
			destVal = reflect.New(destType)
		}
		if err := m.Map(srcVal.Interface(), destVal.Interface()); err != nil {
			return reflect.Value{}, err
		}
		// Match pointer/non-pointer
		if destType.Kind() != reflect.Pointer {
			destVal = destVal.Elem()
		}
		return destVal, nil
	}

	return reflect.Value{}, fmt.Errorf("conversion from %v to %v not implemented", srcType.String(), destType.String())
}
