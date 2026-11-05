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

type dummyVisitor struct{}

func (v *dummyVisitor) VisitField(path string, value any) error {
	return nil
}

type testStructWithSliceOfSlice struct {
	AllowedCACerts [][]byte
}

func TestVisitFields_SliceOfSlice(t *testing.T) {
	obj := &testStructWithSliceOfSlice{
		AllowedCACerts: [][]byte{
			[]byte("cert1"),
			[]byte("cert2"),
		},
	}
	if err := VisitFields(obj, &dummyVisitor{}); err != nil {
		t.Fatalf("VisitFields failed on slice of slice: %v", err)
	}

	emptyObj := &testStructWithSliceOfSlice{}
	if err := VisitFields(emptyObj, &dummyVisitor{}); err != nil {
		t.Fatalf("VisitFields failed on empty slice of slice: %v", err)
	}
}
