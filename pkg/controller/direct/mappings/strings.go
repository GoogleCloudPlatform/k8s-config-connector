package mappings

import (
	"fmt"
	"reflect"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type GoToProtoString struct {
	jsonName string
}

func (m *GoToProtoString) Apply(inRaw any, outRaw any) error {
	in := inRaw.(reflect.Value)
	out := outRaw.(protoreflect.Message)

	if in.Kind() == reflect.Pointer {
		if in.IsNil() {
			return nil
		}
		in = in.Elem()
	}

	_, inFieldIndex, err := findField(in, m.jsonName)
	if err != nil {
		return err
	}

	inValue := in.Field(inFieldIndex)
	if inValue.Kind() == reflect.Ptr {
		if inValue.IsZero() {
			// Nothing to set
			return nil
		}
		inValue = inValue.Elem()
	}
	var s string
	switch inValue.Kind() {
	case reflect.String:
		s = inValue.String()

	default:
		return fmt.Errorf("unexpected kind for field %q, got %v, want string", m.jsonName, inValue.Kind())
	}
	fd := out.Descriptor().Fields().ByJSONName(m.jsonName)
	if fd == nil {
		return fmt.Errorf("proto field with json name %q not known", m.jsonName)
	}
	out.Set(fd, protoreflect.ValueOfString(s))
	// if err := out.ProtoReflect().Set(fd, protoreflect.ValueOfString(s)); err != nil {
	// 	return fmt.Errorf("error setting field %q: %w", m.Field, err)
	// }
	return nil
}

type ProtoToGoString struct {
	jsonName string
}

func (m *ProtoToGoString) Apply(inRaw any, outRaw any) error {
	in := inRaw.(protoreflect.Message)
	out := outRaw.(reflect.Value)

	if out.Kind() == reflect.Ptr {
		out = out.Elem()
	}
	inField := in.Descriptor().Fields().ByJSONName(m.jsonName)
	if inField == nil {
		return fmt.Errorf("proto field with json name %q not known", m.jsonName)
	}

	outField, outFieldIndex, err := findField(out, m.jsonName)
	if err != nil {
		return err
	}

	protoVal := in.Get(inField)
	switch inField.Kind() {
	case protoreflect.StringKind:
		s := protoVal.String()
		if outField.Type.Kind() == reflect.Pointer {
			out.Field(outFieldIndex).Set(reflect.ValueOf(&s))
		} else {
			out.Field(outFieldIndex).Set(reflect.ValueOf(s))
		}
	default:
		return fmt.Errorf("unexpected kind for field %q, got %v, want string", m.jsonName, inField.Kind())
	}
	return nil
}
