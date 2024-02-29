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

package testcmp

import (
	"strings"
	"testing"
)

func UnorderedLineByLineComparisonIgnoreBlankLines(t *testing.T, expected, actual string) {
	expectedLines := make(map[string]int)
	for _, s := range strings.Split(expected, "\n") {
		if s == "" {
			continue
		}
		expectedLines[s]++
	}
	actualLines := strings.Split(actual, "\n")
	for linenum, line := range actualLines {
		if line == "" {
			continue
		}
		v, ok := expectedLines[line]
		if !ok {
			t.Fatalf("actual value didn't match expected value: line %d, %q, not in expected.", linenum, line)
		}
		if v == 1 {
			delete(expectedLines, line)
		} else {
			expectedLines[line] = v - 1
		}
	}
	if len(expectedLines) != 0 {
		t.Fatalf("actual value didn't match expected value: line '%v' not in actual value", expectedLines)
	}
}
