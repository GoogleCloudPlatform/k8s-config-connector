// Copyright 2026 Google LLC
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

package common

import (
	"testing"
)

type sampleDeepStruct struct {
	Value *string
}

type sampleSubStruct struct {
	FieldA *string
	FieldB *int32
	Deep   *sampleDeepStruct
}

type sampleItem struct {
	ID *string
}

type sampleStruct struct {
	Name    *string
	Count   int
	Sub     *sampleSubStruct
	Omitted *sampleSubStruct
	Items   []sampleItem
	Labels  map[string]string
}

func ptrTo[T any](v T) *T {
	return &v
}

func TestCollectPresentFields(t *testing.T) {
	tests := []struct {
		name         string
		structObj    any
		fieldToCheck string
		wantPresent  bool
	}{
		{
			name:         "nil struct - top-level string pointer field is absent",
			structObj:    nil,
			fieldToCheck: "Name",
			wantPresent:  false,
		},
		{
			name:         "empty struct - top-level string pointer field is absent",
			structObj:    &sampleStruct{},
			fieldToCheck: "Name",
			wantPresent:  false,
		},
		{
			name:         "empty struct - top-level non-pointer primitive field is present",
			structObj:    &sampleStruct{},
			fieldToCheck: "Count",
			wantPresent:  true,
		},
		{
			name: "populated top-level string pointer field is present",
			structObj: &sampleStruct{
				Name: ptrTo("hello"),
			},
			fieldToCheck: "Name",
			wantPresent:  true,
		},
		{
			name: "unspecified top-level struct pointer field is absent",
			structObj: &sampleStruct{
				Name: ptrTo("hello"),
			},
			fieldToCheck: "Sub",
			wantPresent:  false,
		},
		{
			name: "populated top-level struct pointer field is present",
			structObj: &sampleStruct{
				Sub: &sampleSubStruct{
					FieldB: ptrTo(int32(42)),
				},
			},
			fieldToCheck: "Sub",
			wantPresent:  true,
		},
		{
			name: "populated nested subfield is present",
			structObj: &sampleStruct{
				Sub: &sampleSubStruct{
					FieldB: ptrTo(int32(42)),
				},
			},
			fieldToCheck: "Sub.FieldB",
			wantPresent:  true,
		},
		{
			name: "unspecified nested subfield is absent",
			structObj: &sampleStruct{
				Sub: &sampleSubStruct{
					FieldB: ptrTo(int32(42)),
				},
			},
			fieldToCheck: "Sub.FieldA",
			wantPresent:  false,
		},
		{
			name: "nil parent struct causes nested path to be absent",
			structObj: &sampleStruct{
				Sub: &sampleSubStruct{},
			},
			fieldToCheck: "Omitted.FieldA",
			wantPresent:  false,
		},
		{
			name: "deeply nested struct field is present",
			structObj: &sampleStruct{
				Sub: &sampleSubStruct{
					Deep: &sampleDeepStruct{
						Value: ptrTo("nested-value"),
					},
				},
			},
			fieldToCheck: "Sub.Deep.Value",
			wantPresent:  true,
		},
		{
			name: "slice of structs elements are present",
			structObj: &sampleStruct{
				Items: []sampleItem{
					{ID: ptrTo("item-1")},
				},
			},
			fieldToCheck: "Items[].ID",
			wantPresent:  true,
		},
		{
			name: "map keys are present",
			structObj: &sampleStruct{
				Labels: map[string]string{
					"env": "production",
				},
			},
			fieldToCheck: "Labels.env",
			wantPresent:  true,
		},
		{
			name: "unspecified map key is absent",
			structObj: &sampleStruct{
				Labels: map[string]string{
					"env": "production",
				},
			},
			fieldToCheck: "Labels.team",
			wantPresent:  false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			present := CollectPresentFields(tc.structObj)
			got := present.Has(tc.fieldToCheck)
			if got != tc.wantPresent {
				t.Errorf("CollectPresentFields() presence for %q = %v, want %v", tc.fieldToCheck, got, tc.wantPresent)
			}
		})
	}
}
