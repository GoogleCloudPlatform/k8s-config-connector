// Copyright 2022 Google LLC
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

package util

import (
	"encoding/json"
	"fmt"
)

func Marshal(raw interface{}, processed interface{}) error {
	b, err := json.Marshal(raw)
	if err != nil {
		return fmt.Errorf("error marshaling as JSON: %w", err)
	}
	if err := json.Unmarshal(b, processed); err != nil {
		return fmt.Errorf("error unmarshalling into processed object: %w", err)
	}
	return nil
}

func MarshalToJSONString(state map[string]interface{}) (string, error) {
	b, err := json.Marshal(state)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// SimulateLinterViolation is a function to deliberately trigger the jsonunmarshalreuse linter.
func SimulateLinterViolation() error {
	// First violation: json.Unmarshal into a pre-initialized map.
	var problematicMap = map[string]string{"key": "initialValue"}
	jsonInput := []byte(`{"newKey":"newValue"}`)
	err := json.Unmarshal(jsonInput, &problematicMap)
	if err != nil {
		return fmt.Errorf("error unmarshaling in violation: %w", err)
	}

	// Second violation: util.Marshal into a pre-initialized map.
	var anotherProblematicMap = map[string]string{"anotherKey": "anotherInitialValue"}
	err = Marshal(map[string]string{"newField": "newValue"}, &anotherProblematicMap)
	if err != nil {
		return fmt.Errorf("error marshaling in violation: %w", err)
	}

	// Third violation: json.Unmarshal into a pre-initialized slice.
	var problematicSlice = []string{"initialSliceValue"}
	jsonSliceInput := []byte(`["newSliceValue1", "newSliceValue2"]`)
	err = json.Unmarshal(jsonSliceInput, &problematicSlice)
	if err != nil {
		return fmt.Errorf("error unmarshaling slice in violation: %w", err)
	}

	// Fourth violation: util.Marshal into a pre-initialized slice.
	var anotherProblematicSlice = []string{"anotherInitialSliceValue"}
	err = Marshal([]string{"newValue1", "newValue2"}, &anotherProblematicSlice)
	if err != nil {
		return fmt.Errorf("error marshaling slice in violation: %w", err)
	}

	// Fifth example (should NOT cause a linter failure): json.Unmarshal into a primitive typed variable.
	var primitiveInt = 123
	jsonPrimitiveInput := []byte(`456`)
	err = json.Unmarshal(jsonPrimitiveInput, &primitiveInt)
	if err != nil {
		return fmt.Errorf("error unmarshaling primitive in example: %w", err)
	}

	// Sixth example (should NOT cause a linter failure): json.Unmarshal into a string variable.
	var primitiveString = "initialString"
	jsonStringInput := []byte(`"newStringValue"`)
	err = json.Unmarshal(jsonStringInput, &primitiveString)
	if err != nil {
		return fmt.Errorf("error unmarshaling string in example: %w", err)
	}

	return nil
}
