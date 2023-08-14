package mappings

import (
	"fmt"
	"reflect"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type Builder interface {
}

type fieldBuilder struct {
	jsonName string
}

func (b *fieldBuilder) Build(in any, out any) (Mapping, error) {
	inNorm, _ := keyFor(in)
	outNorm, _ := keyFor(out)

	switch inNorm := inNorm.(type) {
	case reflect.Value:
		switch outNorm := outNorm.(type) {
		case protoreflect.Message:
			return b.buildGoToProto(inNorm, outNorm)

		default:
			return nil, fmt.Errorf("unhandled type pair %T (recognized as reflect.Value) -> %T (normalized to %T)", in, out, outNorm)
		}
	case protoreflect.Message:
		switch outNorm := outNorm.(type) {
		case reflect.Value:
			return b.buildProtoToGo(inNorm, outNorm)

		default:
			return nil, fmt.Errorf("unhandled type pair %T (recognized as reflect.Value) -> %T (normalized to %T)", in, out, outNorm)
		}
	}

	return nil, fmt.Errorf("unhandled type pair %T (normalized to %T) -> %T (normalized to %T)", in, inNorm, out, outNorm)
}

func findField(in reflect.Value, jsonName string) (*reflect.StructField, int, error) {
	t := in.Type()
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	for i := 0; i < t.NumField(); i++ {
		typeField := t.Field(i)
		if getJSONName(&typeField) == jsonName {
			return &typeField, i, nil
		}
	}
	return nil, 0, fmt.Errorf("field %q not found in %T", jsonName, in)
}

func (b *fieldBuilder) buildGoToProto(in reflect.Value, out protoreflect.Message) (Mapping, error) {
	field, _, err := findField(in, b.jsonName)
	if err != nil {
		return nil, err
	}
	fieldType := field.Type
	if fieldType.Kind() == reflect.Ptr {
		fieldType = fieldType.Elem()
	}
	if fieldType.Kind() == reflect.Slice {
		return &GoToProtoSlice{jsonName: b.jsonName}, nil
	}
	switch fieldType.Kind() {
	case reflect.String:
		return &GoToProtoString{jsonName: b.jsonName}, nil
	case reflect.Struct:
		return &GoToProtoObject{jsonName: b.jsonName}, nil
	}
	return nil, fmt.Errorf("unhandled go-to-proto type %v", fieldType)
}

func (b *fieldBuilder) buildProtoToGo(in protoreflect.Message, out reflect.Value) (Mapping, error) {
	fd := in.Descriptor().Fields().ByJSONName(b.jsonName)
	if fd == nil {
		return nil, fmt.Errorf("unable to find json field %q in message %q", b.jsonName, in.Descriptor().FullName())
	}

	if fd.IsList() {
		return &ProtoToGoSlice{jsonName: b.jsonName}, nil
	}

	fdKind := fd.Kind()
	switch fdKind {
	case protoreflect.StringKind:
		return &ProtoToGoString{jsonName: b.jsonName}, nil
	case protoreflect.MessageKind:
		return &ProtoToGoObject{jsonName: b.jsonName}, nil
	}
	return nil, fmt.Errorf("unhandled proto-to-go type %v", fdKind)
}

func Simple(jsonName string) *fieldBuilder {
	return &fieldBuilder{jsonName: jsonName}
}

func Add(from any, to any, options ...Builder) {
	var mappings []Mapping
	for _, option := range options {
		var mapping Mapping
		switch option := option.(type) {
		case *fieldBuilder:
			m, err := option.Build(from, to)
			if err != nil {
				panic(fmt.Errorf("error building mapper: %v", err))
			}
			mapping = m
		}
		mappings = append(mappings, mapping)
	}

	globalRegistry.Add(from, to, &GoToProtoStruct{Fields: mappings})
}
