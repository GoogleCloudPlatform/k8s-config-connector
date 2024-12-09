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

package fuzz

import (
	"math/rand"
	"reflect"
	"strconv"
	"testing"
)

type RandomFiller struct {
	randStream *rand.Rand

	// for iota based enums, defines the upper bound for a named enum type
	intEnumAllowableValues map[string]int64
	// for non iota based enums, holds the set of allowable values for a named enum type
	stringEnumAllowableValues map[string][]interface{}
}

func NewRandomFiller(seed int64, enumBoundsMap map[string]int64, enumValuesMap map[string][]interface{}) *RandomFiller {
	return &RandomFiller{
		randStream:                rand.New(rand.NewSource(seed)),
		intEnumAllowableValues:    enumBoundsMap, // [0,7]
		stringEnumAllowableValues: enumValuesMap, // {"foo", "bar"}
		// []string{"foo", "bar"} => treat these string fields as ints and fill them with in convertable values
		// another edge case time encoded in a string field
	}
}

// Fill populates the fields of a struct with random values. Enums are handled separately in the
// two maps passed to the RandomFiller.
func (rf *RandomFiller) Fill(t *testing.T, obj interface{}) {
	rf.fillWithRandom(t, reflect.ValueOf(obj).Elem())
}

func (rf *RandomFiller) fillWithRandom(t *testing.T, field reflect.Value) {
	if field.Kind() == reflect.Ptr {
		if field.IsNil() {
			field.Set(reflect.New(field.Type().Elem()))
		}
		rf.fillWithRandom(t, field.Elem())
		return
	}

	switch field.Kind() {
	case reflect.Bool:
		field.SetBool(rf.randStream.Intn(2) == 1)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		// if this field is an iota enum field with a set of allowable values
		if upperBound, ok := rf.intEnumAllowableValues[field.Type().Name()]; ok {
			// Select a random integer value within the range [0, upperBound] for integer enums
			field.SetInt(rf.randStream.Int63n(upperBound + 1))
		} else {
			// Otherwise, fill with a generic random integer
			field.SetInt(rf.randStream.Int63())
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		field.SetUint(rf.randStream.Uint64())

	case reflect.Float32, reflect.Float64:
		field.SetFloat(rf.randStream.Float64())

	case reflect.String:
		// if this field is a enum field with a set of allowable values
		if values, ok := rf.stringEnumAllowableValues[field.Type().Name()]; ok {
			// Select a random string from predefined values in enumValuesMap
			selectedValue := values[rf.randStream.Intn(len(values))]
			field.SetString(selectedValue.(string))
		} else {
			// Otherwise, fill with a generic random string
			field.SetString(randomString(rf.randStream))
		}

	case reflect.Slice:
		count := rf.randStream.Intn(10) + 1
		slice := reflect.MakeSlice(field.Type(), count, count)
		for j := 0; j < count; j++ {
			element := reflect.New(field.Type().Elem()).Elem()
			rf.fillWithRandom(t, element)
			slice.Index(j).Set(element)
		}
		field.Set(slice)

	case reflect.Map:
		count := rf.randStream.Intn(10) + 1
		mapType := reflect.MakeMap(field.Type())
		for j := 0; j < count; j++ {
			key := reflect.New(field.Type().Key()).Elem()
			value := reflect.New(field.Type().Elem()).Elem()
			rf.fillWithRandom(t, key)
			rf.fillWithRandom(t, value)
			mapType.SetMapIndex(key, value)
		}
		field.Set(mapType)

	case reflect.Struct:
		for i := 0; i < field.NumField(); i++ {
			//t.Log("todo acpana struct field names:", field.Type().Field(i).Name )
			if field.Type().Field(i).Name == "SyncWaitSecs" { // this func capitlizes the first letter
				structField := field.Field(i)
				// TODO double check
				if structField.Kind() == reflect.Ptr {
					if structField.IsNil() {
						structField.Set(reflect.New(structField.Type().Elem()))
					}
					structField = structField.Elem()
					// rf.fillWithRandom(t, field.Elem())
					// return
				}
				structField.SetString(strconv.FormatInt(rf.randStream.Int63(), 10))
				t.Log("todo acpana it worked")
				continue
			}
			if field.Type().Field(i).Name == "AuditIntervalSeconds" { // this func capitlizes the first letter
				structField := field.Field(i)
				//t.Logf("todo acpana before operation -- Field: %s, Type: %s", structField.Type().Name(), structField.Type().Elem())
				// TODO double check
				if structField.Kind() == reflect.Ptr {
					t.Logf("todo acpana before operation -- Field: %s, Type: %s", structField.Type().Name(), structField.Type().Elem())
					if structField.IsNil() {
						structField.Set(reflect.New(structField.Type().Elem()))
					}
					structField = structField.Elem()
					// rf.fillWithRandom(t, field.Elem())
					// return
				} else {
					t.Logf("todo acpana before operation -- Field: %s, Type: %s", structField.Type().Name(), structField.Type())
				}
				t.Logf("todo acpana after operation -- Field: %s, Type: %s", structField.Type().Name(), structField.Type())

				// there are two AuditIntervalSeconds fields: one *string one *int64
				switch structField.Kind() {
				case reflect.String:
					structField.SetString(strconv.FormatInt(rf.randStream.Int63(), 10))
				case reflect.Int64:
					structField.SetInt(rf.randStream.Int63())
				default:
					t.Errorf("AuditIntervalSeconds type undefined: %v", structField.Kind())
				}

				continue
			}
			rf.fillWithRandom(t, field.Field(i))
		}

	default:
		t.Fatalf("Unhandled field kind: %v", field.Kind())
	}
}
