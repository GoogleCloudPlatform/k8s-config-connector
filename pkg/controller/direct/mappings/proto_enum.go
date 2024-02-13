package mappings

// import (
// 	"reflect"

// 	"google.golang.org/protobuf/reflect/protoreflect"
// 	"k8s.io/klog/v2"
// )

// type enumToStringTypeMapping struct {
// 	// scope *Mapping

// 	enumType       *reflectType
// 	enumDescriptor protoreflect.EnumDescriptor
// 	stringType     *reflectType
// }

// var _ TypeMapping = &enumToStringTypeMapping{}

// func (m *enumToStringTypeMapping) FromType() reflect.Type {
// 	return m.enumType.rt
// }

// func (m *enumToStringTypeMapping) ToType() reflect.Type {
// 	return m.stringType.rt
// }

// func (m *enumToStringTypeMapping) Map(in *point, out *point) error {
// 	enumVal := in.GetValue()
// 	klog.Fatalf("type is %T", enumVal.Interface())
// 	// protoIn := protoreflect.ValueOf(inVal.Interface())
// 	return nil
// }

// type stringToEnumTypeMapping struct {
// 	// scope *Mapping

// 	enumType       *reflectType
// 	enumDescriptor protoreflect.EnumDescriptor
// 	stringType     *reflectType
// }

// var _ TypeMapping = &stringToEnumTypeMapping{}

// func (m *stringToEnumTypeMapping) FromType() reflect.Type {
// 	return m.stringType.rt
// }

// func (m *stringToEnumTypeMapping) ToType() reflect.Type {
// 	return m.enumType.rt
// }

// func (m *stringToEnumTypeMapping) Map(in *point, out *point) error {
// 	stringVal := in.GetValue().Interface().(string)
// 	klog.Infof("stringVal is %+v", stringVal)
// 	// protoIn := protoreflect.ValueOf(inVal.Interface())
// 	enumValues := m.enumDescriptor.Values()
// 	for i := 0; i < enumValues.Len(); i++ {
// 		enumValue := enumValues.Get(i)
// 		klog.Infof("name is %+v", enumValue.Name())
// 		if string(enumValue.Name()) == stringVal {
// 			out.SetValue(enumValue)
// 			return nil
// 		}
// 	}

// 	// klog.Fatalf("type is %T", stringVal.Interface())
// 	return nil
// }
