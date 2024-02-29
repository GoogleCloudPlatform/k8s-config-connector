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

package deepcopy

func StringStringMap(src map[string]string) map[string]string {
	if src == nil {
		return nil
	}
	dst := make(map[string]string, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

// DeepCopy deeply copies map[string]interface{}, []interface{}, and primitives types.
func DeepCopy(o interface{}) interface{} {
	if m, ok := o.(map[string]interface{}); ok {
		newO := make(map[string]interface{})
		for k, v := range m {
			newO[k] = DeepCopy(v)
		}
		return newO
	}
	if s, ok := o.([]interface{}); ok {
		newO := make([]interface{}, 0)
		for _, v := range s {
			newO = append(newO, DeepCopy(v))
		}
		return newO
	}
	return o
}

func MapStringInterface(o map[string]interface{}) map[string]interface{} {
	return DeepCopy(o).(map[string]interface{})
}

func StringSlice(src []string) []string {
	dst := make([]string, len(src))
	for i, v := range src {
		dst[i] = v
	}
	return dst
}
