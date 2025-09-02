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

package test

import (
	"fmt"
	"strings"
)

// ApplyReplacementsToMap applies `replacements“ for all the values in the map (recursively).
func ApplyReplacementsToMap(m map[string]any, replacements map[string]string) map[string]any {
	for k, v := range m {
		v = applyReplacementsToAny(v, replacements)
		m[k] = v
	}
	return m
}

func applyReplacementsToAny(v any, replacements map[string]string) any {
	switch v := v.(type) {
	case string:
		return applyReplacementsToString(v, replacements)

	case map[string]any:
		return ApplyReplacementsToMap(v, replacements)
	case []any:
		return applyReplacementsToSlice(v, replacements)

	case nil:
		return v

	case int64:
		return v

	case bool:
		return v

	default:
		panic(fmt.Sprintf("unhandled type %T", v))
	}
}

func applyReplacementsToString(s string, replacements map[string]string) string {
	for k, v := range replacements {
		s = strings.ReplaceAll(s, k, v)
	}
	return s
}

func applyReplacementsToSlice(s []any, replacements map[string]string) []any {
	for i, v := range s {
		v = applyReplacementsToAny(v, replacements)
		s[i] = v
	}
	return s
}
