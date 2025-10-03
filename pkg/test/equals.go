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

package test

import (
	"encoding/json"
	"reflect"
	"testing"
)

func Equals(t *testing.T, o1, o2 interface{}) bool {
	return reflect.DeepEqual(normalize(t, o1), normalize(t, o2))
}

func normalize(t *testing.T, obj interface{}) map[string]interface{} {
	// Only applicable to types that can me marshaled as a JSON object.
	b, err := json.Marshal(&obj)
	if err != nil {
		t.Fatalf("error marshaling as JSON: %v", err)
	}
	var ret map[string]interface{}
	if err := json.Unmarshal(b, &ret); err != nil {
		t.Fatalf("error unmarshalling JSON bytes: %v", err)
	}
	return ret
}
