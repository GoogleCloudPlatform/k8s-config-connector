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

package monitoring

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/testing/protocmp"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

func lastComponent(s string) string {
	i := strings.LastIndex(s, "/")
	return s[i+1:]
}

func ComputeChangedFields(actual proto.Message, desired proto.Message) sets.Set[string] {
	changes := sets.New[string]()
	actualReflect := actual.ProtoReflect()
	desiredReflect := desired.ProtoReflect()
	actualReflect.Range(func(field protoreflect.FieldDescriptor, actualValue protoreflect.Value) bool {
		desiredValue := desiredReflect.Get(field)
		if !actualValue.Equal(desiredValue) {
			changes.Insert(field.JSONName())
		}
		return true
	})
	desiredReflect.Range(func(field protoreflect.FieldDescriptor, desiredValue protoreflect.Value) bool {
		actualValue := actualReflect.Get(field)
		if !actualValue.Equal(desiredValue) {
			changes.Insert(field.JSONName())
		}
		return true
	})
	if changes.Len() != 0 {
		klog.V(2).Infof("ComputeChangedFields found diff fields=%v, diff=%v", sets.List(changes), cmp.Diff(actual, desired, protocmp.Transform()))
	}
	return changes
}

func setStatus(u *unstructured.Unstructured, typedStatus any) error {
	status, err := runtime.DefaultUnstructuredConverter.ToUnstructured(typedStatus)
	if err != nil {
		return fmt.Errorf("error converting status to unstructured: %w", err)
	}

	// Use existing values for conditions/observedGeneration; they are managed in k8s not the GCP API
	old, _, _ := unstructured.NestedMap(u.Object, "status")
	if old != nil {
		status["conditions"] = old["conditions"]
		status["observedGeneration"] = old["observedGeneration"]
	}

	u.Object["status"] = status

	return nil
}

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
				w.visitAny(path+"[]", v.Index(i))
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
