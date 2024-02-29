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
	"fmt"
	"reflect"

	"k8s.io/klog/v2"
)

type point struct {
	scope *Mapping
	rv    reflect.Value
	t     *reflectType
}

func (m *Mapping) newPoint(rv reflect.Value) *point {
	t := typeOf(rv.Type())

	return &point{
		scope: m,
		rv:    rv,
		t:     t,
	}
}

func (p *point) Child(id FieldID) *point {
	if p == nil {
		return nil
	}
	field := p.t.findField(id)
	if field == nil {
		klog.Warningf("unable to find field %q in %v", id, p.t)
		return nil
	}
	return field.getValue(p)
}

func (p *point) SetValue(field FieldID, value reflect.Value) error {
	f := p.t.findField(field)
	if f == nil {
		return fmt.Errorf("unable to find field %q in %v", field, p.t)
	}

	return f.setValue(p, value)
}

func (p *point) GetValue() reflect.Value {
	return p.rv
}
