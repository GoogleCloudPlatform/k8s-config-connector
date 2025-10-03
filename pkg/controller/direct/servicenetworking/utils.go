// Copyright 2025 Google LLC
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

package servicenetworking

import (
	"encoding/json"

	"k8s.io/klog/v2"
)

// ReflectClone creates a deep copy of the given object.
// It is implemented by serializing and unserializing the object to/from json.
func ReflectClone[T any](obj T) T {
	j, err := json.Marshal(obj)
	if err != nil {
		klog.Fatalf("failed to marshal object to json: %v", err)
	}
	var cloneObj T
	if err := json.Unmarshal(j, &cloneObj); err != nil {
		klog.Fatalf("failed to unmarshal json to object: %v", err)
	}
	return cloneObj
}
