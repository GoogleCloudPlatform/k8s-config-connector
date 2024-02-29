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

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// protoOneOfField is a Field that is backed by a oneof value in proto.
type protoOneOfField struct {
	t       *reflectType
	jsonKey string
	fd      protoreflect.FieldDescriptor
}

var _ Field = &protoOneOfField{}

// ID implements the Field interface.
func (f *protoOneOfField) ID() FieldID {
	fieldID := toFieldID(f.jsonKey)
	return fieldID
}

// getJSONKey implements the Field interface.
func (f *protoOneOfField) getJSONKey() string {
	return f.jsonKey
}

// Type implements the Field interface.
func (f *protoOneOfField) Type() *reflectType {
	return f.t
}

// isRequired implements the Field interface.
func (f *protoOneOfField) isRequired() bool {
	return false
}

// getValue implements the Field interface.
func (f *protoOneOfField) getValue(p *point) *point {
	protoMsg := p.rv.Interface().(proto.Message)
	fdVal := protoMsg.ProtoReflect().Get(f.fd)
	if !fdVal.IsValid() {
		return nil
	}
	obj := fdVal.Message().Interface()
	rv := reflect.ValueOf(obj)
	out := p.scope.newPoint(rv)
	return out
}

// setValue implements the Field interface.
func (f *protoOneOfField) setValue(p *point, srcVal reflect.Value) error {
	if srcVal.IsNil() {
		return nil
	}

	protoMsg := p.rv.Interface().(proto.Message)
	fdVal := protoMsg.ProtoReflect().Get(f.fd)

	fdObj := fdVal.Message().New()

	destType := reflect.TypeOf(fdObj.Interface())

	destVal, err := p.scope.convert(srcVal, destType)
	if err != nil {
		return fmt.Errorf("converting %v to %v: %w", srcVal.Type(), destType, err)
	}
	if !destVal.IsValid() {
		return nil
	}

	destMsg := destVal.Interface().(proto.Message)
	protoMsg.ProtoReflect().Set(f.fd, protoreflect.ValueOfMessage(destMsg.ProtoReflect()))

	return nil
}
