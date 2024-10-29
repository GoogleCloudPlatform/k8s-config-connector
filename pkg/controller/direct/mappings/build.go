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

// Mapping holds the logic for mapping a particular resource,
// including sub-object types.
type Mapping struct {
	ResourceCloudType *reflectType
	ResourceKRMType   *reflectType

	Mappings []typeMapping
}

// Map will convert an object from "in" to "out", mapping object subfields recursively.
func (m *Mapping) Map(in any, out any) error {
	inVal := reflect.ValueOf(in)
	outVal := reflect.ValueOf(out)

	inPoint := m.newPoint(inVal)
	outPoint := m.newPoint(outVal)

	for _, typeMapping := range m.Mappings {
		if typeMapping.FromType() != inPoint.t.rt {
			continue
		}

		if typeMapping.ToType() != outPoint.t.rt {
			continue
		}

		return typeMapping.Map(inPoint, outPoint)
	}

	for _, typeMapping := range m.Mappings {
		klog.Infof("mapping from %q -> %q", typeMapping.FromType(), typeMapping.ToType())
	}

	return fmt.Errorf("type mapping not found for %q -> %q", inPoint.t, outPoint.t)
}

// FieldMapping is the base interface for per-field mappings.
type FieldMapping interface {
	// TODO: We should have something here, because right now everything is a FieldMapping
	// BuildFieldMapping(m *TypeMapping) (*fieldMapping, error)
}

// statusField describes a mapping from a proto top-level field into a KRM status field.
type statusField struct {
	ID string
}

// Status fields are a proto top-level field but nested under Status in KRM.
func Status(id string) FieldMapping {
	return &statusField{ID: id}
}

// statusField describes a mapping from a proto top-level field into a KRM spec field.
type specField struct {
	ID string
}

// Status fields are a proto top-level field but nested under Spec in KRM.
func Spec(id string) FieldMapping {
	return &specField{ID: id}
}

// ignoreField describes a field that should be actively ignored, we know it should not be automatically mapped.
// The machinery may still process it directly.
type ignoreField struct {
	ID FieldID
}

// Ignore fields should not be mapped automatically.
// They are different from TODO fields, in that we have actively determined they should not be mapped.
func Ignore(id string) FieldMapping {
	return &ignoreField{ID: toFieldID(id)}
}

// MappingBuilder allows for fluid construction of a Mapping
type MappingBuilder struct {
	mapping *Mapping
	errors  []error
}

// Build "finalizes" the mapping, and returns the constructed mapping.
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

// MustBuild is like Build, but panics on error.
func (b *MappingBuilder) MustBuild() *Mapping {
	m, err := b.Build()
	if err != nil {
		klog.Fatalf("error building mapping: %v", err)
	}
	return m
}

// NewMapping starts a new mappingBuilder, for converting proto <-> KRM.
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

	b = b.addKRMToCloudMapping(resourceKRMType, resourceCloudType, true, fields...)
	b = b.addCloudToKRMMapping(resourceCloudType, resourceKRMType, true, fields...)
	return b

}

// MapNested describes how a nested subobject should be mapped, when it is encountered in this context.
func (b *MappingBuilder) MapNested(cloudObj any, krmObj any, fields ...FieldMapping) *MappingBuilder {
	cloudVal := reflect.ValueOf(cloudObj)
	krmVal := reflect.ValueOf(krmObj)
	resourceCloudType := typeOf(cloudVal.Type())
	resourceKRMType := typeOf(krmVal.Type())

	b = b.addKRMToCloudMapping(resourceKRMType, resourceCloudType, false, fields...)
	b = b.addCloudToKRMMapping(resourceCloudType, resourceKRMType, false, fields...)
	return b
}

// addKRMToCloudMapping will add a mapping for mapping from KRM to cloud objects.
func (b *MappingBuilder) addKRMToCloudMapping(inType *reflectType, outType *reflectType, hasSpecStatus bool, fields ...FieldMapping) *MappingBuilder {
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
				InPath:  parseFieldPath("spec." + field.ID),
				OutPath: parseFieldPath(field.ID),
			})

		case *statusField:
			createMapping.fields = append(createMapping.fields, &fieldMapping{
				InPath:  parseFieldPath("status." + field.ID),
				OutPath: parseFieldPath(field.ID),
			})

			// For simple 1:1 mappings, we can just accept a string
		case string:
			createMapping.fields = append(createMapping.fields, &fieldMapping{
				InPath:  parseFieldPath(field),
				OutPath: parseFieldPath(field),
			})

		case *ignoreField:
			createMapping.ignore = append(createMapping.ignore, field)

		default:
			klog.Fatalf("unhandled field type %T", field)
		}
	}

	return b
}

// addCloudToKRMMapping will add a mapping for mapping from Cloud to KRM objects.
func (b *MappingBuilder) addCloudToKRMMapping(inType *reflectType, outType *reflectType, hasSpecStatus bool, fields ...FieldMapping) *MappingBuilder {
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
				InPath:  parseFieldPath(field.ID),
				OutPath: parseFieldPath("spec." + field.ID),
			})

		case *statusField:
			createMapping.fields = append(createMapping.fields, &fieldMapping{
				InPath:  parseFieldPath(field.ID),
				OutPath: parseFieldPath("status." + field.ID),
			})

			// For simple 1:1 mappings, we can just accept a string
		case string:
			createMapping.fields = append(createMapping.fields, &fieldMapping{
				InPath:  parseFieldPath(field),
				OutPath: parseFieldPath(field),
			})

		case *ignoreField:
			createMapping.ignore = append(createMapping.ignore, field)

		default:
			klog.Fatalf("unhandled field type %T", field)
		}
	}

	return b
}
