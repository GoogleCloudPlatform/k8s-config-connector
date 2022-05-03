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

package pathslice

import "strings"

// Base returns the last element of the path.
// Returns empty string if path is empty.
func Base(path []string) string {
	if len(path) == 0 {
		return ""
	}
	return path[len(path)-1]
}

// ToString returns the path as a string with elements delimited by periods.
func ToString(path []string) string {
	return strings.Join(path, ".")
}
