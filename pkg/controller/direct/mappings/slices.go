package mappings

import (
	"fmt"
	"reflect"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type GoToProtoSlice struct {
	jsonName string
	// ProtobufType proto.Message
}

func (m *GoToProtoSlice) Apply(inRaw any, outRaw any) error {
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
	var inSlice reflect.Value
	switch inValue.Kind() {
	case reflect.Slice:
		inSlice = inValue

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

	switch outField.Kind() {
	case protoreflect.StringKind:
		outSlice := out.NewField(outField).List()
		n := inSlice.Len()
		for i := 0; i < n; i++ {
			inValue := inSlice.Index(i)
			var s string
			switch inValue.Kind() {
			case reflect.String:
				s = inValue.String()
			default:
				return fmt.Errorf("unexpected type for value, got %v, want string", inValue.Kind())
			}
			outSlice.Append(protoreflect.ValueOfString(s))
		}
		out.Set(outField, protoreflect.ValueOfList(outSlice))
	case protoreflect.MessageKind:
		outSlice := out.NewField(outField).List()
		n := inSlice.Len()
		for i := 0; i < n; i++ {
			inValue := inSlice.Index(i)
			outValue := outSlice.AppendMutable()
			if err := globalRegistry.Map(inValue, outValue); err != nil {
				return err
			}
		}

		out.Set(outField, protoreflect.ValueOfList(outSlice))
	default:
		return fmt.Errorf("unhandled type %v for list", outField.Kind())
	}

	return nil
}

type ProtoToGoSlice struct {
	jsonName string
	// ProtobufType proto.Message
}

func (m *ProtoToGoSlice) Apply(inRaw any, outRaw any) error {
	in := inRaw.(protoreflect.Message)
	out := outRaw.(reflect.Value)

	if out.Kind() == reflect.Ptr {
		if out.IsNil() {
			v := reflect.New(out.Type().Elem())
			out.Set(reflect.ValueOf(v.Interface()))
		}
		out = out.Elem()
	}

	inField := in.Descriptor().Fields().ByJSONName(m.jsonName)
	if inField == nil {
		return fmt.Errorf("proto field with json name %q not known", m.jsonName)
	}

	_, outFieldIndex, err := findField(out, m.jsonName)
	if err != nil {
		return err
	}

	outVal := out.Field(outFieldIndex)

	inFieldValue := in.Get(inField)
	inFieldList := inFieldValue.List()

	n := inFieldList.Len()
	for i := 0; i < n; i++ {
		protoVal := inFieldList.Get(i)
		switch inField.Kind() {
		case protoreflect.StringKind:
			s := protoVal.String()
			outVal = reflect.Append(outVal, reflect.ValueOf(s))
		case protoreflect.MessageKind:
			outItem := reflect.New(outVal.Type().Elem()).Elem()
			if err := Map(protoVal, outItem); err != nil {
				return err
			}
			outVal = reflect.Append(outVal, outItem)
		default:
			return fmt.Errorf("unexpected kind for field %q, got %v, want string", m.jsonName, inField.Kind())
		}
	}

	out.Field(outFieldIndex).Set(outVal)

	return nil
}
