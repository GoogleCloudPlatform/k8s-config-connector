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

// Field represents a field in an object type.
type Field interface {
	// ID returns the unique ID of the field, typically the value used in JSON serialization.
	ID() FieldID

	// Type returns the type of the field, using our reflectType mapper.
	Type() *reflectType

	// getValue returns a handle to the field value of the specified parent.
	getValue(parent *point) *point

	// setValue sets
	setValue(parent *point, src reflect.Value) error

	// getJSONKey returns the key we will normally use for JSON serialization.
	// This is normally the same as the ID value.
	// This is used for suggesting a go field declaration in validation.
	getJSONKey() string

	// isRequired returns true if the field is required in JSON.
	// This is used for suggesting a go field declaration in validation.
	isRequired() bool
}

// FieldID is the unique identifier for a field within an object type.
type FieldID string

// toFieldID converts a string to a FieldID.
func toFieldID(id string) FieldID {
	// TODO: Do we want to normalize this here?
	// return FieldID(strings.ToLower(id))
	return FieldID(id)
}
