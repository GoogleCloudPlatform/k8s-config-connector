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

type TypeMapping interface {
	FromType() reflect.Type
	ToType() reflect.Type

	Map(in *point, out *point, inFieldPrefix *fieldPath) error
}

type structTypeMapping struct {
	scope *Mapping

	inType  *reflectType
	outType *reflectType

	hasSpecStatus bool

	fields []*fieldMapping

	ignore []*ignoreField
}

var _ TypeMapping = &structTypeMapping{}

func (m *structTypeMapping) FromType() reflect.Type {
	return m.inType.rt
}

func (m *structTypeMapping) ToType() reflect.Type {
	return m.outType.rt
}

type fieldMapping struct {
	InPath    *fieldPath
	OutPath   *fieldPath
	Transform func(in reflect.Value) (reflect.Value, error)
}

func (m *structTypeMapping) Map(in *point, out *point, inFieldPrefix *fieldPath) error {
	for _, mapping := range m.fields {
		if !mapping.InPath.HasPrefix(inFieldPrefix) {
			continue
		}

		inPoint := mapping.InPath.FindPoint(in)

		srcVal := inPoint.GetValue()
		if mapping.Transform != nil {
			xformed, err := mapping.Transform(srcVal)
			if err != nil {
				return err
			}
			srcVal = xformed
		}

		if err := mapping.OutPath.SetValue(out, srcVal); err != nil {
			return err
		}

	}

	return nil
}
