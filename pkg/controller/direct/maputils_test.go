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

package direct

import (
	"encoding/json"
	"testing"

	kccutil "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util"
	"google.golang.org/protobuf/types/known/durationpb"
)

func TestStringDuration_FromProto(t *testing.T) {
	mapctx := &MapContext{}
	d := &durationpb.Duration{Seconds: 34312, Nanos: 20}
	krm := StringDuration_FromProto(mapctx, d)
	if *krm != "9h31m52.00000002s" {
		t.Fatalf("google.protobuf.Duration -> string, expect \"9h31m52.00000002s\", got %s", *krm)
	}
	if mapctx.Err() != nil {
		t.Fatalf("google.protobuf.Duration -> string error: %s", mapctx.Err())
	}
}

func TestStringDuration_ToProto(t *testing.T) {
	mapctx := &MapContext{}
	s := "1h1m"
	d := StringDuration_ToProto(mapctx, &s)
	if d.Seconds != 3660 || d.Nanos != 0 {
		t.Fatalf("string -> google.protobuf.Duration, expect \"seconds:3660 nanos:00\", got %s", d)
	}
	if mapctx.Err() != nil {
		t.Fatalf("google.protobuf.Duration -> String error: %s", mapctx.Err())
	}
}

// TestSimulateLinterViolationInDirect is a test function to deliberately trigger the jsonunmarshalreuse linter in the direct package.
func TestSimulateLinterViolationInDirect(t *testing.T) {
	// 1. json.Unmarshal into a pre-initialized map.
	var problematicMap = map[string]string{"key": "initialValue"}
	jsonInput := []byte(`{"newKey":"newValue"}`)
	err := json.Unmarshal(jsonInput, &problematicMap)
	if err != nil {
		t.Errorf("error unmarshaling in violation 1: %v", err)
	}

	// 2. util.Marshal into a pre-initialized map.
	var anotherProblematicMap = map[string]string{"anotherKey": "anotherInitialValue"}
	err = kccutil.Marshal(map[string]string{"newField": "newValue"}, &anotherProblematicMap)
	if err != nil {
		t.Errorf("error marshaling in violation 2: %v", err)
	}

	// 3. json.Unmarshal into a pre-initialized slice.
	var problematicSlice = []string{"initialSliceValue"}
	jsonSliceInput := []byte(`["newSliceValue1", "newSliceValue2"]`)
	err = json.Unmarshal(jsonSliceInput, &problematicSlice)
	if err != nil {
		t.Errorf("error unmarshaling slice in violation 3: %v", err)
	}

	// 4. util.Marshal into a pre-initialized slice.
	var anotherProblematicSlice = []string{"anotherInitialSliceValue"}
	err = kccutil.Marshal([]string{"newValue1", "newValue2"}, &anotherProblematicSlice)
	if err != nil {
		t.Errorf("error marshaling slice in violation 4: %v", err)
	}

	// 5. json.Unmarshal into a primitive typed variable (should NOT cause a linter failure).
	var primitiveInt = 123
	jsonPrimitiveInput := []byte(`456`)
	err = json.Unmarshal(jsonPrimitiveInput, &primitiveInt)
	if err != nil {
		t.Errorf("error unmarshaling primitive in violation 5: %v", err)
	}

	// 6. json.Unmarshal into a string variable (should NOT cause a linter failure).
	var primitiveString = "initialString"
	jsonStringInput := []byte(`"newStringValue"`)
	err = json.Unmarshal(jsonStringInput, &primitiveString)
	if err != nil {
		t.Errorf("error unmarshaling string in violation 6: %v", err)
	}
}
