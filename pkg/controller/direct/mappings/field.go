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

type Field interface {
	ID() FieldID
	JSONKey() string
	Type() *reflectType

	getValue(p *point) *point
	setValue(p *point, src reflect.Value) error

	isRequired() bool
}

type FieldID string

func ToFieldID(id string) FieldID {
	// TODO: Do we want to normalize?
	// return FieldID(strings.ToLower(id))
	return FieldID(id)
}
