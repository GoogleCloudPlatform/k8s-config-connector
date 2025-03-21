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
		if err := w.visitor.VisitField(path, v.Interface()); err != nil {
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
		case reflect.Struct, reflect.String:
			for i := 0; i < v.Len(); i++ {
				// We pass the address, so that the visitor can mutate the value in place
				w.visitAny(path+"[]", v.Index(i).Addr())
			}
		case reflect.Uint8:
			// Do not visit []byte as individual values, treat as a leaf
		default:
			w.errs = append(w.errs, fmt.Errorf("visiting slice of type %v is not supported", elemType.Kind()))
		}

	case reflect.String, reflect.Bool, reflect.Int32, reflect.Int64, reflect.Float64:
		// "leaf", nothing to recurse into
	default:
		w.errs = append(w.errs, fmt.Errorf("visiting type %v is not supported", v.Kind()))
	}
}
