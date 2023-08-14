package mappings

import (
	"reflect"
	"strings"
)

// func mapCloudToKube(in *pb.Key, out *v1alpha1.APIKeysKeySpec) error {
// 	mapping, err := mappings.Get(in, out)
// 	if err != nil {
// 		return err
// 	}
// 	return mapping.Apply(in.ProtoReflect(), reflect.ValueOf(out))
// }

type Mapping interface {
	Apply(in any, out any) error
}

func getJSONName(f *reflect.StructField) string {
	jsonTag := f.Tag.Get("json")
	switch jsonTag {
	case "-":
		// Not marshaled
		return ""
	case "":
		return f.Name

	default:
		parts := strings.Split(jsonTag, ",")
		name := parts[0]
		if name == "" {
			name = f.Name
		}
		return name
	}
}

type GoToProtoStruct struct {
	Fields []Mapping
}

func (m *GoToProtoStruct) Apply(in any, out any) error {
	for _, mapping := range m.Fields {
		if err := mapping.Apply(in, out); err != nil {
			return err
		}
	}
	return nil
}

type ProtoToGoStruct struct {
	Fields []Mapping
}

func (m *ProtoToGoStruct) Apply(in any, out any) error {
	for _, mapping := range m.Fields {
		if err := mapping.Apply(in, out); err != nil {
			return err
		}
	}
	return nil
}
