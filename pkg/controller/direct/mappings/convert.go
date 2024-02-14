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

func (m *Mapping) convert(src reflect.Value, destType reflect.Type) (reflect.Value, error) {
	if src.Type() == destType {
		// Nothing to do
		return src, nil
	}

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
		// klog.Infof("copied slice %v -> %v", src.Interface(), dest.Interface())
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
		// case "*durationpb.Duration":
		// 	// TODO: Register some well-known conversion functions and use them?
		// 	duration, err := time.ParseDuration(v)
		// 	if err != nil {
		// 		return reflect.Value{}, fmt.Errorf("invalid duration %q", duration)
		// 	}
		// 	durationProto := durationpb.New(duration)
		// 	return reflect.ValueOf(durationProto), nil
		// case "*v1alpha1.ResourceRef":
		// 	// TODO: Register some well-known conversion functions and use them?
		// 	ref := &v1alpha1.ResourceRef{External: v}
		// 	return reflect.ValueOf(ref), nil
		// case "v1alpha1.ResourceRef":
		// 	// TODO: Register some well-known conversion functions and use them?
		// 	ref := &v1alpha1.ResourceRef{External: v}
		// 	return reflect.ValueOf(ref).Elem(), nil
		default:
			return reflect.Value{}, fmt.Errorf("string conversion to %v not implemented", destType.String())
		}
	// case "durationpb.Duration":
	// 	v := src.Addr().Interface().(*durationpb.Duration)
	// 	switch destType.String() {
	// 	case "string":
	// 		s := v.AsDuration().String()
	// 		return reflect.ValueOf(s), nil
	// 	case "*string":
	// 		s := v.AsDuration().String()
	// 		return reflect.ValueOf(&s), nil
	// 	default:
	// 		return reflect.Value{}, fmt.Errorf("duration conversion to %v not implemented", destType.String())
	// 	}
	// case "v1alpha1.ResourceRef":
	// 	v := src.Addr().Interface().(*v1alpha1.ResourceRef)
	// 	switch destType.String() {
	// 	case "string":
	// 		s := v.External
	// 		return reflect.ValueOf(s), nil
	// 	case "*string":
	// 		s := v.External
	// 		return reflect.ValueOf(&s), nil
	// 	default:
	// 		return reflect.Value{}, fmt.Errorf("v1alpha1.ResourceRef conversion to %v not implemented", destType.String())
	// 	}
	case "int":
		v64 := src.Int()
		switch destType.String() {
		case "int64":
			// TODO: int64 <-> int should actually be a validation warning, we're using an ill-defined type in KRM
			return reflect.ValueOf(v64), nil
		case "int32":
			// TODO: int32 <-> int should actually be a validation warning, we're using an ill-defined type in KRM
			i32 := int32(v64)
			return reflect.ValueOf(i32), nil
		case "*int32":
			// TODO: int32 <-> int should actually be a validation warning, we're using an ill-defined type in KRM
			i32 := int32(v64)
			return reflect.ValueOf(&i32), nil
		default:
			return reflect.Value{}, fmt.Errorf("int conversion to %v not implemented", destType.String())
		}
	case "int64":
		v64 := src.Int()
		switch destType.String() {
		case "int":
			// TODO: int64 <-> int should actually be a validation warning, we're using an ill-defined type in KRM
			v := int(v64)
			return reflect.ValueOf(v), nil
		case "*int":
			v := int(v64)
			return reflect.ValueOf(&v), nil
		default:
			return reflect.Value{}, fmt.Errorf("int64 conversion to %v not implemented", destType.String())
		}
	case "int32":
		v32 := int32(src.Int())
		switch destType.String() {
		case "int":
			// TODO: int32 <-> int should actually be a validation warning, we're using an ill-defined type in KRM
			v := int(v32)
			return reflect.ValueOf(v), nil
		case "*int":
			v := int(v32)
			return reflect.ValueOf(&v), nil
		default:
			return reflect.Value{}, fmt.Errorf("int64 conversion to %v not implemented", destType.String())
		}
	}

	klog.Warningf("fallthrough on src %q", srcType.String())

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
