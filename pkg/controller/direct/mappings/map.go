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

import "reflect"

// typeMapping is a mapping from one type to another.
type typeMapping interface {
	// FromType returns the type we will convert from
	FromType() reflect.Type
	// ToType returns the type we will convert to.
	ToType() reflect.Type

	// Map performs the actual conversion from one value to another.
	Map(in *point, out *point) error
}

// structTypeMapping is a typeMapping for mapping between struct fields
type structTypeMapping struct {
	//scope is the parent Mapping context of which we are part.
	scope *Mapping

	// inType is the type we will convert from.
	inType *reflectType
	// outType is the type we will convert to.
	outType *reflectType

	// hasSpecStatus is a hint for validation that this mapping is a top-level mapping.
	hasSpecStatus bool

	// fields defines the mapping of fields
	fields []*fieldMapping

	// ignore is a list of fields we should ignore.
	// This is kept around for validation.
	ignore []*ignoreField
}

var _ typeMapping = &structTypeMapping{}

// FromType implements typeMapping
func (m *structTypeMapping) FromType() reflect.Type {
	return m.inType.rt
}

// ToType implements typeMapping
func (m *structTypeMapping) ToType() reflect.Type {
	return m.outType.rt
}

// fieldMapping describes the field we should read in the input, and write in the output.
// In future, we could add value transformations here, but we haven't needed one yet.
type fieldMapping struct {
	InPath  *fieldPath
	OutPath *fieldPath
}

// Map implements typeMapping
func (m *structTypeMapping) Map(in *point, out *point) error {
	for _, mapping := range m.fields {
		inPoint := mapping.InPath.FindPoint(in)

		srcVal := inPoint.GetValue()

		if err := mapping.OutPath.SetValue(out, srcVal); err != nil {
			return err
		}

	}

	return nil
}
