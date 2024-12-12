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
	"encoding/json"
	"math/rand"
	"reflect"
	"strconv"
	"testing"
)

// This is a placeholder struct for all possible types, native or not,
// that can be found in a KRM struct.
type TestKRMType struct {
	BoolField          bool
	BoolFieldPtr       *bool
	IntField           int
	IntFieldPtr        *int
	StringField        string
	StringFieldPtr     *string
	FloatField         float64
	FloatFieldPtr      *float64
	SliceField         []string
	SliceFieldPtr      *[]string
	MapField           map[string]int
	MapFieldPtr        *map[string]int
	ComplexMapField    map[string]NestedKRMType
	ComplexMapFieldPtr *map[string]NestedKRMType
	NestedStruct       NestedKRMType
	NestedStructPtr    *NestedKRMType
	// EnumField          MyEnumType
	// EnumFieldPtr       *MyEnumType
	// StringEnumField    MyStringEnumType
	// StringEnumFieldPtr *MyStringEnumType
	StructSlice    []NestedKRMType
	StructSlicePtr *[]NestedKRMType
	IntSlice       []int
	IntSlicePtr    *[]int

	PtrStringFieldThatNeedsInt *string
	StringFieldThatNeedsInt    string
}

// Nested structs
type NestedKRMType struct {
	NestedIntField     int
	NestedIntFieldPtr  *int
	NestedBoolField    bool
	NestedBoolFieldPtr *bool
	DeepNestedField    *DeepNestedKRMType
}

type DeepNestedKRMType struct {
	FieldA    string
	FieldAPtr *string
	FieldB    float64
	FieldBPtr *float64
	FieldC    []byte
	FieldCPtr *[]byte
}

// // Integer enum type
// type MyEnumType int

// const (
// 	EnumValueA MyEnumType = iota
// 	EnumValueB
// 	EnumValueC
// )

// // String enum type
// type MyStringEnumType string

// const (
// 	StringEnumOptionA MyStringEnumType = "OptionA"
// 	StringEnumOptionB MyStringEnumType = "OptionB"
// 	StringEnumOptionC MyStringEnumType = "OptionC"
// )

func TestRandomFillerFields(t *testing.T) {
	stream := rand.New(rand.NewSource(int64(9201995))) // for determinism

	funcF := func(t *testing.T, fieldName string, field reflect.Value) {
		field.SetString(strconv.FormatInt(stream.Int63(), 10))
	}

	overrides := map[string]OverrideFiller{
		".PtrStringFieldThatNeedsInt": funcF,
		".StringFieldThatNeedsInt":    funcF,
	}
	filler := NewRandomFiller(&FillerConfig{Stream: stream, FieldOverrides: overrides})

	tests := []struct {
		name       string
		fieldCheck func(krmObj *TestKRMType) bool
	}{
		// Not all field types will have a valid test for us to check with: See BoolField.

		//{"BoolField", func(krmObj *TestKRMType) bool { return krmObj.BoolField }}, // Both False and True are valid.
		{"BoolFieldPtr", func(krmObj *TestKRMType) bool { return krmObj.BoolFieldPtr != nil }},
		{"IntField", func(krmObj *TestKRMType) bool { return krmObj.IntField != 0 }},
		{"IntFieldPtr", func(krmObj *TestKRMType) bool { return krmObj.IntFieldPtr != nil && *krmObj.IntFieldPtr != 0 }},
		{"StringField", func(krmObj *TestKRMType) bool { return krmObj.StringField != "" }},
		{"StringFieldPtr", func(krmObj *TestKRMType) bool { return krmObj.StringFieldPtr != nil && *krmObj.StringFieldPtr != "" }},
		{"FloatField", func(krmObj *TestKRMType) bool { return krmObj.FloatField != 0 }},
		{"FloatFieldPtr", func(krmObj *TestKRMType) bool { return krmObj.FloatFieldPtr != nil && *krmObj.FloatFieldPtr != 0 }},
		{"SliceField", func(krmObj *TestKRMType) bool { return len(krmObj.SliceField) > 0 }},
		{"SliceFieldPtr", func(krmObj *TestKRMType) bool { return krmObj.SliceFieldPtr != nil && len(*krmObj.SliceFieldPtr) > 0 }},
		{"MapField", func(krmObj *TestKRMType) bool { return len(krmObj.MapField) > 0 }},
		{"MapFieldPtr", func(krmObj *TestKRMType) bool { return krmObj.MapFieldPtr != nil && len(*krmObj.MapFieldPtr) > 0 }},
		{"ComplexMapField", func(krmObj *TestKRMType) bool { return len(krmObj.ComplexMapField) > 0 }},
		{"ComplexMapFieldPtr", func(krmObj *TestKRMType) bool {
			return krmObj.ComplexMapFieldPtr != nil && len(*krmObj.ComplexMapFieldPtr) > 0
		}},
		// {"EnumField", func(krmObj *TestKRMType) bool { return krmObj.EnumField >= 0 && krmObj.EnumField <= 2 }},
		// {"EnumFieldPtr", func(krmObj *TestKRMType) bool {
		// 	return krmObj.EnumFieldPtr != nil && *krmObj.EnumFieldPtr >= 0 && *krmObj.EnumFieldPtr <= 2
		// }},
		// {"StringEnumField", func(krmObj *TestKRMType) bool { return krmObj.StringEnumField != "" }},
		// {"StringEnumFieldPtr", func(krmObj *TestKRMType) bool {
		// 	return krmObj.StringEnumFieldPtr != nil && *krmObj.StringEnumFieldPtr != ""
		// }},
		{"StructSlice", func(krmObj *TestKRMType) bool { return len(krmObj.StructSlice) > 0 }},
		{"StructSlicePtr", func(krmObj *TestKRMType) bool { return krmObj.StructSlicePtr != nil && len(*krmObj.StructSlicePtr) > 0 }},
		{"IntSlice", func(krmObj *TestKRMType) bool { return len(krmObj.IntSlice) > 0 }},
		{"IntSlicePtr", func(krmObj *TestKRMType) bool { return krmObj.IntSlicePtr != nil && len(*krmObj.IntSlicePtr) > 0 }},
		{"NestedStruct.NestedIntField", func(krmObj *TestKRMType) bool { return krmObj.NestedStruct.NestedIntField != 0 }},
		{"NestedStruct.NestedIntFieldPtr", func(krmObj *TestKRMType) bool {
			return krmObj.NestedStruct.NestedIntFieldPtr != nil && *krmObj.NestedStruct.NestedIntFieldPtr != 0
		}},
		//{"NestedStruct.NestedBoolField", func(krmObj *TestKRMType) bool { return krmObj.NestedStruct.NestedBoolField }},
		{"NestedStruct.NestedBoolFieldPtr", func(krmObj *TestKRMType) bool { return krmObj.NestedStruct.NestedBoolFieldPtr != nil }},
		{"NestedStruct.DeepNestedField.FieldA", func(krmObj *TestKRMType) bool {
			return krmObj.NestedStruct.DeepNestedField != nil && krmObj.NestedStruct.DeepNestedField.FieldA != ""
		}},
		{"NestedStruct.DeepNestedField.FieldAPtr", func(krmObj *TestKRMType) bool {
			return krmObj.NestedStruct.DeepNestedField != nil && krmObj.NestedStruct.DeepNestedField.FieldAPtr != nil && *krmObj.NestedStruct.DeepNestedField.FieldAPtr != ""
		}},
		{"NestedStruct.DeepNestedField.FieldB", func(krmObj *TestKRMType) bool {
			return krmObj.NestedStruct.DeepNestedField != nil && krmObj.NestedStruct.DeepNestedField.FieldB != 0
		}},
		{"NestedStruct.DeepNestedField.FieldBPtr", func(krmObj *TestKRMType) bool {
			return krmObj.NestedStruct.DeepNestedField != nil && krmObj.NestedStruct.DeepNestedField.FieldBPtr != nil && *krmObj.NestedStruct.DeepNestedField.FieldBPtr != 0
		}},
		{"NestedStruct.DeepNestedField.FieldC", func(krmObj *TestKRMType) bool {
			return krmObj.NestedStruct.DeepNestedField != nil && len(krmObj.NestedStruct.DeepNestedField.FieldC) > 0
		}},
		{"NestedStruct.DeepNestedField.FieldCPtr", func(krmObj *TestKRMType) bool {
			return krmObj.NestedStruct.DeepNestedField != nil && krmObj.NestedStruct.DeepNestedField.FieldCPtr != nil && len(*krmObj.NestedStruct.DeepNestedField.FieldCPtr) > 0
		}},
		{"PtrStringFieldThatNeedsInt", func(krmObj *TestKRMType) bool {
			if krmObj.PtrStringFieldThatNeedsInt == nil {
				return false
			}
			if _, err := strconv.ParseInt(*krmObj.PtrStringFieldThatNeedsInt, 10, 64); err != nil {
				t.Error("converting string field to int")
			}
			return true
		}},
		{"StringFieldThatNeedsInt", func(krmObj *TestKRMType) bool {
			if krmObj.PtrStringFieldThatNeedsInt == nil {
				return false
			}
			if _, err := strconv.ParseInt(*krmObj.PtrStringFieldThatNeedsInt, 10, 64); err != nil {
				t.Error("converting string field to int")
			}
			return true
		}},
	}

	krmObj := &TestKRMType{}
	filler.Fill(t, krmObj)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.fieldCheck(krmObj) {
				jsonData, err := json.MarshalIndent(krmObj, "", "  ")
				if err != nil {
					t.Error(err)
				}

				t.Fatalf("Field %s was not filled as expected; struct: %s", tt.name, string(jsonData))
			}
		})
	}
}
