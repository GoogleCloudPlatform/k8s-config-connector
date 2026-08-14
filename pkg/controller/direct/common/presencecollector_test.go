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

type sampleSubStruct struct {
	FieldA *string
	FieldB *int32
}

type sampleStruct struct {
	Name    *string
	Sub     *sampleSubStruct
	Omitted *sampleSubStruct
}

func TestCollectPresentFields(t *testing.T) {
	strVal := "hello"
	intVal := int32(42)

	obj := &sampleStruct{
		Name: &strVal,
		Sub: &sampleSubStruct{
			FieldB: &intVal,
			// FieldA is nil
		},
		// Omitted is nil
	}

	present := CollectPresentFields(obj)

	if !present.Has("Name") {
		t.Errorf("expected 'Name' to be present")
	}
	if !present.Has("Sub") {
		t.Errorf("expected 'Sub' to be present")
	}
	if !present.Has("Sub.FieldB") {
		t.Errorf("expected 'Sub.FieldB' to be present")
	}
	if present.Has("Sub.FieldA") {
		t.Errorf("expected 'Sub.FieldA' to NOT be present")
	}
	if present.Has("Omitted") {
		t.Errorf("expected 'Omitted' to NOT be present")
	}
	if present.Has("Omitted.FieldA") {
		t.Errorf("expected 'Omitted.FieldA' to NOT be present")
	}
}
