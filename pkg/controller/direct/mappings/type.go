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
	"sync"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"k8s.io/klog/v2"
)

type reflectType struct {
	rt reflect.Type

	fields map[FieldID]Field
}

type typeCache struct {
	mutex sync.Mutex
	types map[reflect.Type]*reflectType
}

var allTypes typeCache

func (c *typeCache) get(t reflect.Type) *reflectType {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.types == nil {
		c.types = make(map[reflect.Type]*reflectType)
	}

	if tt, ok := c.types[t]; ok {
		return tt
	}

	tt := &reflectType{rt: t}
	c.types[t] = tt
	return tt
}

func typeOf(t reflect.Type) *reflectType {
	return allTypes.get(t)
}

func (t *reflectType) String() string {
	prefix := ""
	reflectType := t.rt

	for {
		if reflectType.Kind() == reflect.Ptr {
			prefix = "*" + prefix
			reflectType = reflectType.Elem()
		} else if reflectType.Kind() == reflect.Slice {
			prefix = "[]" + prefix
			reflectType = reflectType.Elem()
		} else {
			break
		}
	}

	return prefix + reflectType.PkgPath() + "." + reflectType.Name()
}

func (t *reflectType) AssignableTo(interfaceType *reflectType) bool {
	return t.rt.AssignableTo(interfaceType.rt)
}

func (t *reflectType) lookupField(fields *fieldPath) Field {
	n := len(fields.fields)
	p := t.findField(fields.fields[0])
	if p == nil {
		return nil
	}
	for i := 1; i < n; i++ {
		f := p.Type().findField(fields.fields[i])
		if f == nil {
			return nil
		}
		p = f
	}
	return p
}

func (t *reflectType) Fields() map[FieldID]Field {
	if t.fields != nil {
		return t.fields
	}
	var fieldList []Field
	reflectType := t.rt

	protoMessageInterfaceType := reflect.TypeOf((*proto.Message)(nil)).Elem()

	// proto oneof (union) fields have a complicated structure,
	// but can easily be discovered and get/set through protoreflect
	if reflectType.Implements(protoMessageInterfaceType) {
		obj := reflect.New(reflectType).Elem().Interface().(proto.Message)

		descriptor := obj.ProtoReflect().Descriptor()

		fields := descriptor.Fields()
		for i := 0; i < fields.Len(); i++ {
			fd := fields.Get(i)

			oneOf := fd.ContainingOneof()
			if oneOf == nil {
				continue
			}

			fdObj := obj.ProtoReflect().Get(fd)

			// klog.Infof("found oneof %s", oneOf)

			field := &protoOneOfField{fd: fd}
			field.jsonKey = fd.JSONName()
			switch fd.Kind() {
			case protoreflect.MessageKind:
				t := reflect.TypeOf(fdObj.Message().Interface())
				field.t = allTypes.get(t)
			default:
				klog.Fatalf("cannot handle oneof member field %v", fd)
			}

			fieldList = append(fieldList, field)
		}
	}

	if reflectType.Kind() == reflect.Ptr {
		reflectType = reflectType.Elem()
	}
	n := reflectType.NumField()
	for i := 0; i < n; i++ {
		rf := reflectType.Field(i)
		if !rf.IsExported() {
			continue
		}
		// Skip proto oneof (as we handled them above)
		oneOfTag := rf.Tag.Get("protobuf_oneof")
		if oneOfTag != "" {
			continue
		}
		f := &structField{f: &rf, jsonKey: getJSONFieldTag(&rf)}
		fieldList = append(fieldList, f)
	}

	fields := make(map[FieldID]Field)
	for _, f := range fieldList {
		id := f.ID()
		if fields[id] != nil {
			klog.Fatalf("duplicate field %s", f.ID())
		}
		fields[id] = f
	}

	t.fields = fields
	return fields
}

func (t *reflectType) findField(id FieldID) Field {
	f := t.Fields()[id]
	return f
}
