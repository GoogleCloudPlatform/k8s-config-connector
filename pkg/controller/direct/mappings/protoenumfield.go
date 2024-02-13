package mappings

import (
	"fmt"
	"reflect"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"k8s.io/klog/v2"
)

type protoEnumField struct {
	t       *reflectType
	jsonKey string
	fd      protoreflect.FieldDescriptor
	enum    protoreflect.EnumDescriptor
}

var _ Field = &protoEnumField{}

func (f *protoEnumField) ID() FieldID {
	fieldID := ToFieldID(f.jsonKey)
	return fieldID
}

// JSONKey returns the key we will normally use for JSON serialization.
func (f *protoEnumField) JSONKey() string {
	return f.jsonKey
}

func (f *protoEnumField) Type() *reflectType {
	return f.t
}

func (f *protoEnumField) isRequired() bool {
	return false
}

func (f *protoEnumField) getValue(p *point) *point {
	protoMsg := p.rv.Interface().(proto.Message)
	fdVal := protoMsg.ProtoReflect().Get(f.fd)
	if !fdVal.IsValid() {
		return nil
	}
	enumValueDescriptor := f.enum.Values().ByNumber(fdVal.Enum())
	if enumValueDescriptor == nil {
		klog.Warningf("enum value %v is not known", fdVal)
		return nil
	}
	rv := reflect.ValueOf(string(enumValueDescriptor.Name()))
	out := p.scope.newPoint(rv)
	return out
}

func (f *protoEnumField) setValue(p *point, srcVal reflect.Value) error {

	var stringValue string
	if srcVal.Kind() == reflect.Pointer {
		if srcVal.IsNil() {
			// Maybe set to undefined?
			return nil
		}
		srcVal = srcVal.Elem()
	}

	switch srcVal.Kind() {
	case reflect.String:
		stringValue = srcVal.String()
	default:
		return fmt.Errorf("unsupported value kind when setting enum: %s", srcVal.Kind())
	}

	protoMsg := p.rv.Interface().(proto.Message)

	enumValues := f.enum.Values()
	for i := 0; i < enumValues.Len(); i++ {
		enumValue := enumValues.Get(i)
		if string(enumValue.Name()) == stringValue {
			protoMsg.ProtoReflect().Set(f.fd, protoreflect.ValueOf(enumValue.Number()))
			return nil
		}
	}

	return fmt.Errorf("unknown enum value %q", stringValue)
}
