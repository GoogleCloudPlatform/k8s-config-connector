// Copyright 2025 Google LLC
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

package tags

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type FieldChange struct {
	FieldPath    string
	ActualValue  protoreflect.Value
	DesiredValue protoreflect.Value
}

func fieldHasChanged(fieldPath string, desired protoreflect.Message, actual protoreflect.Message) (*FieldChange, error) {
	change := &FieldChange{FieldPath: fieldPath}

	actualValue, foundActual, err := commonGetFieldByPath(actual, fieldPath)
	if err != nil {
		return change, err
	}
	change.ActualValue = actualValue

	desiredValue, foundDesired, err := commonGetFieldByPath(desired, fieldPath)
	if err != nil {
		return change, err
	}
	change.DesiredValue = desiredValue

	if foundActual != foundDesired {
		return change, nil
	}
	if !foundActual && !foundDesired {
		// Both unset
		return change, nil
	}
	if actualValue.Equal(desiredValue) {
		// Note: returning nil to indicate no change
		return nil, nil
	}
	return change, nil
}

func commonGetFieldByPath(msg protoreflect.Message, fieldPath string) (protoreflect.Value, bool, error) {
	if msg == nil {
		return protoreflect.Value{}, false, nil
	}
	tokens := strings.SplitN(fieldPath, ".", 2)
	fieldName := protoreflect.Name(tokens[0])
	field := msg.Descriptor().Fields().ByName(fieldName)
	if field == nil {
		return protoreflect.Value{}, false, fmt.Errorf("field %q not found in %T", fieldName, msg)
	}
	v := msg.Get(field)
	if len(tokens) == 1 {
		return v, true, nil
	}
	switch field.Kind() {
	case protoreflect.MessageKind:
		return commonGetFieldByPath(v.Message(), tokens[1])
	default:
		return protoreflect.Value{}, false, fmt.Errorf("field %q in %T is not a message", fieldName, msg)
	}
}

func format(v protoreflect.Value) string {
	o := v.Interface()
	if msg, ok := o.(protoreflect.Message); ok {
		return prototext.Format(msg.Interface())
	}
	return fmt.Sprintf("[%T]:%v", o, o)
}
