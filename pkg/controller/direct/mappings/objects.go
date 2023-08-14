package mappings

import (
	"fmt"
	"reflect"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type GoToProtoObject struct {
	jsonName string
	// ProtobufType proto.Message
}

func (m *GoToProtoObject) Apply(inRaw any, outRaw any) error {
	in := inRaw.(reflect.Value)
	out := outRaw.(protoreflect.Message)

	if in.Kind() == reflect.Ptr {
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
		if inValue.IsNil() {
			return nil
		}
		inValue = inValue.Elem()
	}
	var inValueStruct any
	switch inValue.Kind() {
	case reflect.Struct:
		inValueStruct = inValue

	default:
		return fmt.Errorf("unexpected kind for field %q, got %v, want object", m.jsonName, inValue.Kind())
	}
	// inMap, ok := inAny.(map[string]interface{})
	// if !ok {
	// 	return fmt.Errorf("unexpected type for field %q, got %T, want object", m.Field, inAny)
	// }
	outField := out.Descriptor().Fields().ByJSONName(m.jsonName)
	if outField == nil {
		return fmt.Errorf("proto field with json name %q not known", m.jsonName)
	}

	outValue := out.NewField(outField)
	if err := globalRegistry.Map(inValueStruct, outValue); err != nil {
		return err
	}

	out.Set(outField, outValue)

	return nil
}

type ProtoToGoObject struct {
	jsonName string
}

func (m *ProtoToGoObject) Apply(inRaw any, outRaw any) error {
	in := inRaw.(protoreflect.Message)
	out := outRaw.(reflect.Value)

	inField := in.Descriptor().Fields().ByJSONName(m.jsonName)
	if inField == nil {
		return fmt.Errorf("proto field with json name %q not known", m.jsonName)
	}

	_, outFieldIndex, err := findField(out, m.jsonName)
	if err != nil {
		return err
	}

	protoVal := in.Get(inField)
	switch inField.Kind() {
	case protoreflect.MessageKind:
		s := protoVal.Message()
		outVal := out.Field(outFieldIndex)
		if err := Map(s, outVal); err != nil {
			return err
		}
		out.Field(outFieldIndex).Set(outVal)
	default:
		return fmt.Errorf("unexpected kind for field %q, got %v", m.jsonName, inField.Kind())
	}
	return nil
}
