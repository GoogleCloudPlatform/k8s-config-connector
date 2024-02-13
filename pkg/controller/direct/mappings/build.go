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
	"errors"
	"fmt"
	"os"
	"reflect"

	"k8s.io/klog/v2"
)

type Mapping struct {
	ResourceCloudType *reflectType
	ResourceKRMType   *reflectType

	Mappings []TypeMapping
}

func (m *Mapping) Map(in any, out any) error {
	// inVal := &ReflectValue{reflect.ValueOf(in)}
	// outVal := &ReflectValue{reflect.ValueOf(out)}

	inVal := reflect.ValueOf(in)
	outVal := reflect.ValueOf(out)

	// inType := typeOf(inVal.Type())
	// outType := typeOf(outVal.Type())

	inPoint := m.newPoint(inVal)
	outPoint := m.newPoint(outVal)

	// TODO: Create Elem or Deinterface method?
	// if outType.rt.Kind() == reflect.Pointer {
	// 	if outType.rt.Elem().Kind() == reflect.Interface {
	// 		outType = typeOf(reflect.ValueOf(out).Elem())
	// 	}
	// }

	// inTypeKey := inType.String()
	// outTypeKey := outType.String()
	for _, typeMapping := range m.Mappings {
		if typeMapping.FromType() != inPoint.t.rt {
			continue
		}

		if typeMapping.ToType() != outPoint.t.rt {
			continue
		}

		return typeMapping.Map(inPoint, outPoint)
	}

	// // Check for interfaces
	// if inType.AssignableTo(outType) {
	// 	inVal := &ReflectValue{reflect.ValueOf(in)}
	// 	outVal := &ReflectValue{reflect.ValueOf(out)}
	// 	klog.Infof("interface assignable; setting %v to %v", outVal.Type(), inVal.Type())
	// 	klog.Infof("interface assignable; setting %+v to %+v", outVal.rv.Interface(), inVal.rv.Interface())
	// 	outVal.rv.Elem().Set(inVal.rv)
	// 	return nil
	// }

	for _, typeMapping := range m.Mappings {
		klog.Infof("mapping from %q -> %q", typeMapping.FromType(), typeMapping.ToType())
	}

	return fmt.Errorf("type mapping not found for %q -> %q", inPoint.t, outPoint.t)
}

type FieldMapping interface {
	// BuildFieldMapping(m *TypeMapping) (*fieldMapping, error)
}

type statusField struct {
	ID string
}

func Status(id string) FieldMapping {
	return &statusField{ID: id}
}

type ignoreField struct {
	ID FieldID
}

func Ignore(id string) FieldMapping {
	return &ignoreField{ID: ToFieldID(id)}
}

func TODO(id string) FieldMapping {
	return &ignoreField{ID: ToFieldID(id)}
}

type resourceID struct {
	ID string
}

func ResourceID(id string) FieldMapping {
	return &resourceID{ID: id}
}

type specField struct {
	ID string
}

func Spec(id string) FieldMapping {
	return &specField{ID: id}
}

type resourceRef struct {
	ID     string
	Mapper ResourceRefMapper
}

type ResourceRefMapper = Mapper

func ResourceRef(id string, mapper ResourceRefMapper) FieldMapping {
	return &resourceRef{ID: id, Mapper: mapper}
}

type transformedFieldMapping struct {
	ID     string
	Mapper ResourceRefMapper
}

type Mapper interface {
	KRMToCloud(in reflect.Value) (reflect.Value, error)
	CloudToKRM(in reflect.Value) (reflect.Value, error)
}

func Transformed(id string, mapper Mapper) FieldMapping {
	return &transformedFieldMapping{ID: id, Mapper: mapper}
}

type MappingBuilder struct {
	mapping *Mapping
	errors  []error
}

func (b *MappingBuilder) Build() (*Mapping, error) {
	if len(b.errors) != 0 {
		return nil, errors.Join(b.errors...)
	}

	if AssertEnabled {
		errs := b.mapping.Validate()
		if len(errs) != 0 {
			for _, err := range errs {
				fmt.Fprintf(os.Stderr, "%v\n", err.Message)
				if err.Proposal != "" {
					fmt.Fprintf(os.Stderr, "    %v\n", err.Proposal)
				}
			}
			AssertFail()
		}
	}

	return b.mapping, nil
}

func (b *MappingBuilder) MustBuild() *Mapping {
	m, err := b.Build()
	if err != nil {
		klog.Fatalf("error building mapping: %v", err)
	}
	return m
}

func NewMapping(cloudObj any, krmObj any, fields ...FieldMapping) *MappingBuilder {
	cloudVal := reflect.ValueOf(cloudObj)
	krmVal := reflect.ValueOf(krmObj)
	resourceCloudType := typeOf(cloudVal.Type())
	resourceKRMType := typeOf(krmVal.Type())

	m := &Mapping{
		ResourceCloudType: resourceCloudType,
		ResourceKRMType:   resourceKRMType,
	}

	b := &MappingBuilder{
		mapping: m,
	}

	b = b.mapKRMToCloud(resourceKRMType, resourceCloudType, true, fields...)
	b = b.mapCloudToKRM(resourceCloudType, resourceKRMType, true, fields...)
	return b

}

// func (b *MappingBuilder) WithCreate(out any, fields ...Field) *MappingBuilder {
// 	return b.mapKRMToCloud(b.mapping.ResourceKRMType, typeOf(out), fields...)
// }

// func (b *MappingBuilder) WithDelete(out any, fields ...Field) *MappingBuilder {
// 	return b.mapKRMToCloud(b.mapping.ResourceKRMType, typeOf(out), fields...)
// }

func (b *MappingBuilder) MapNested(cloudObj any, krmObj any, fields ...FieldMapping) *MappingBuilder {
	cloudVal := reflect.ValueOf(cloudObj)
	krmVal := reflect.ValueOf(krmObj)
	resourceCloudType := typeOf(cloudVal.Type())
	resourceKRMType := typeOf(krmVal.Type())

	b = b.mapKRMToCloud(resourceKRMType, resourceCloudType, false, fields...)
	b = b.mapCloudToKRM(resourceCloudType, resourceKRMType, false, fields...)
	return b
}

func (b *MappingBuilder) mapKRMToCloud(inType *reflectType, outType *reflectType, hasSpecStatus bool, fields ...FieldMapping) *MappingBuilder {

	createMapping := &structTypeMapping{
		scope:         b.mapping,
		inType:        inType,
		outType:       outType,
		hasSpecStatus: hasSpecStatus,
	}
	b.mapping.Mappings = append(b.mapping.Mappings, createMapping)

	for _, field := range fields {
		switch field := field.(type) {
		case *specField:
			createMapping.fields = append(createMapping.fields, &fieldMapping{
				InPath:  ParseFieldPath("spec." + field.ID),
				OutPath: ParseFieldPath(field.ID),
			})

		case *resourceRef:
			createMapping.fields = append(createMapping.fields, &fieldMapping{
				InPath:    ParseFieldPath(field.ID),
				OutPath:   ParseFieldPath(field.ID),
				Transform: field.Mapper.KRMToCloud,
			})

		case *transformedFieldMapping:
			createMapping.fields = append(createMapping.fields, &fieldMapping{
				InPath:    ParseFieldPath(field.ID),
				OutPath:   ParseFieldPath(field.ID),
				Transform: field.Mapper.KRMToCloud,
			})

		case *statusField:
			createMapping.fields = append(createMapping.fields, &fieldMapping{
				InPath:  ParseFieldPath("status." + field.ID),
				OutPath: ParseFieldPath(field.ID),
			})

		case *resourceID:
			createMapping.fields = append(createMapping.fields, &fieldMapping{
				InPath:  ParseFieldPath("spec.resourceID"),
				OutPath: ParseFieldPath(field.ID),
			})

		case string:
			createMapping.fields = append(createMapping.fields, &fieldMapping{
				InPath:  ParseFieldPath(field),
				OutPath: ParseFieldPath(field),
			})

		case *ignoreField:
			createMapping.ignore = append(createMapping.ignore, field)

		default:
			klog.Fatalf("unhandled field type %T", field)
		}
	}

	return b
}

func (b *MappingBuilder) mapCloudToKRM(inType *reflectType, outType *reflectType, hasSpecStatus bool, fields ...FieldMapping) *MappingBuilder {

	createMapping := &structTypeMapping{
		scope:         b.mapping,
		inType:        inType,
		outType:       outType,
		hasSpecStatus: hasSpecStatus,
	}
	b.mapping.Mappings = append(b.mapping.Mappings, createMapping)

	for _, field := range fields {
		switch field := field.(type) {
		case *specField:
			createMapping.fields = append(createMapping.fields, &fieldMapping{
				InPath:  ParseFieldPath(field.ID),
				OutPath: ParseFieldPath("spec." + field.ID),
			})

		case *resourceRef:
			createMapping.fields = append(createMapping.fields, &fieldMapping{
				InPath:    ParseFieldPath(field.ID),
				OutPath:   ParseFieldPath(field.ID),
				Transform: field.Mapper.CloudToKRM,
			})

		case *transformedFieldMapping:
			createMapping.fields = append(createMapping.fields, &fieldMapping{
				InPath:    ParseFieldPath(field.ID),
				OutPath:   ParseFieldPath(field.ID),
				Transform: field.Mapper.CloudToKRM,
			})

		case *statusField:
			createMapping.fields = append(createMapping.fields, &fieldMapping{
				InPath:  ParseFieldPath(field.ID),
				OutPath: ParseFieldPath("status." + field.ID),
			})

		case *resourceID:
			createMapping.fields = append(createMapping.fields, &fieldMapping{
				InPath:  ParseFieldPath("spec.resourceID"),
				OutPath: ParseFieldPath(field.ID),
			})

		case string:
			createMapping.fields = append(createMapping.fields, &fieldMapping{
				InPath:  ParseFieldPath(field),
				OutPath: ParseFieldPath(field),
			})

		case *ignoreField:
			createMapping.ignore = append(createMapping.ignore, field)

		default:
			klog.Fatalf("unhandled field type %T", field)
		}
	}

	return b
}

// func (b *MappingBuilder) MapEnum(enum protoreflect.EnumType) *MappingBuilder {
// 	enumDescriptor := enum.Descriptor()
// 	cloudObj := enum.New(0)
// 	cloudVal := reflect.ValueOf(cloudObj)
// 	// krmVal := reflect.ValueOf(krmObj)
// 	// cloudType := typeOf(cloudVal.Type())
// 	// resourceKRMType := typeOf(krmVal.Type())
// 	enumType := typeOf(reflect.PtrTo(cloudVal.Type()))
// 	// klog.Fatalf("enumType is %v", enumType)

// 	var krmObj string
// 	krmVal := reflect.ValueOf(krmObj)
// 	stringType := typeOf(krmVal.Type())

// 	{
// 		createMapping := &enumToStringTypeMapping{
// 			// scope:         b.mapping,
// 			enumType:       enumType,
// 			enumDescriptor: enumDescriptor,
// 			stringType:     stringType,
// 			// hasSpecStatus: false,
// 		}
// 		b.mapping.Mappings = append(b.mapping.Mappings, createMapping)
// 	}

// 	{
// 		createMapping := &stringToEnumTypeMapping{
// 			// scope:         b.mapping,
// 			enumType:       enumType,
// 			enumDescriptor: enumDescriptor,
// 			stringType:     stringType,
// 			// hasSpecStatus: false,
// 		}
// 		b.mapping.Mappings = append(b.mapping.Mappings, createMapping)
// 	}

// 	return b
// }
