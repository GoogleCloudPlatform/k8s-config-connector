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

package common

import (
	"errors"
	"fmt"
	"reflect"
)

type Visitor interface {
	VisitField(path string, value any) error
}

func VisitFields(obj any, visitor Visitor) error {
	w := &visitorWalker{visitor: visitor}
	w.visitAny("", reflect.ValueOf(obj))
	return errors.Join(w.errs...)
}

type visitorWalker struct {
	visitor Visitor
	errs    []error
}

func (w *visitorWalker) visitAny(path string, v reflect.Value) {
	shouldCallVisitor := true
	switch v.Kind() {
	case reflect.Ptr:
		if v.IsNil() {
			// Skip nil pointers
			shouldCallVisitor = false
		}
	}
	if shouldCallVisitor {
		val := v.Interface()
		if v.Kind() != reflect.Ptr && v.CanAddr() {
			val = v.Addr().Interface()
		}
		if err := w.visitor.VisitField(path, val); err != nil {
			w.errs = append(w.errs, err)
		}
	}

	switch v.Kind() {
	case reflect.Ptr, reflect.Interface:
		if v.IsNil() {
			return
		}
		w.visitAny(path, v.Elem())

	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			field := v.Type().Field(i)
			if field.IsExported() {
				fieldName := field.Name
				w.visitAny(path+"."+fieldName, v.Field(i))
			}
		}

	case reflect.Map:
		for _, key := range v.MapKeys() {
			w.visitAny(path+"."+key.String(), v.MapIndex(key))
		}

	case reflect.Slice:
		elemType := v.Type().Elem()
		switch elemType.Kind() {
		case reflect.Struct:
			for i := 0; i < v.Len(); i++ {
				elem := v.Index(i)
				if elem.CanAddr() {
					w.visitAny(path+"[]", elem.Addr())
				} else {
					w.visitAny(path+"[]", elem)
				}
			}
		case reflect.Ptr:
			for i := 0; i < v.Len(); i++ {
				elem := v.Index(i)
				if !elem.IsNil() {
					w.visitAny(path+"[]", elem.Elem())
				}
			}
		case reflect.String:
			for i := 0; i < v.Len(); i++ {
				w.visitAny(path+"[]", v.Index(i))
			}
		case reflect.Uint8, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.Bool:
			// Treat primitive slices as leaves
		default:
			w.errs = append(w.errs, fmt.Errorf("visiting slice of type %v is not supported", elemType.Kind()))
		}

	case reflect.String, reflect.Bool, reflect.Int, reflect.Int32, reflect.Int64, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
		// "leaf", nothing to recurse into
	default:
		w.errs = append(w.errs, fmt.Errorf("visiting type %v is not supported", v.Kind()))
	}
}
