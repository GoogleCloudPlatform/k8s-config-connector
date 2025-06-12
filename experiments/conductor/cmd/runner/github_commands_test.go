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

package runner

import (
	"regexp"
	"testing"
)

func TestCommitMessageRegex(t *testing.T) {
	tests := []struct {
		name           string
		regex          *regexp.Regexp
		input          string
		expectedResult bool
	}{
		{
			name:           "Match REGEX_MSG_50",
			regex:          REGEX_MSG_50,
			input:          "TestKind: Move existing test to subdirectory",
			expectedResult: true,
		},
		{
			name:           "No match REGEX_MSG_50",
			regex:          REGEX_MSG_50,
			input:          ": Move existing test to subdirectory",
			expectedResult: false,
		},
		{
			name:           "Match REGEX_MSG_51A",
			regex:          REGEX_MSG_51A,
			input:          "TestKind: Move existing test to subdirectory before creating maximal test",
			expectedResult: true,
		},
		{
			name:           "No match REGEX_MSG_51A",
			regex:          REGEX_MSG_51A,
			input:          ": Move existing test to subdirectory before creating maximal test",
			expectedResult: false,
		},
		{
			name:           "Match REGEX_MSG_51 using COMMIT_MSG_51B",
			regex:          REGEX_MSG_51,
			input:          "TestKind: Create maximal test",
			expectedResult: true,
		},
		{
			name:           "Match REGEX_MSG_51 using COMMIT_MSG_51C",
			regex:          REGEX_MSG_51,
			input:          "TestKind: Enable test with mockgcp",
			expectedResult: true,
		},
		{
			name:  "Match REGEX_MSG_51 using COMMIT_MSG_51B and COMMIT_MSG_51C",
			regex: REGEX_MSG_51,
			// According to the regex `*([a-zA-Z]+: Create maximal test)|([a-zA-Z]+: Enable test with mockgcp).*$`,
			// when there are text after `[a-zA-Z]+: Create maximal test`, it's also a match.
			input:          "TestKind: Create maximal test | TestKind: Enable test with mockgcp",
			expectedResult: true,
		},
		{
			name:           "No match REGEX_MSG_51 using COMMIT_MSG_51A",
			regex:          REGEX_MSG_51,
			input:          "TestKind: Move existing test to subdirectory before creating maximal test",
			expectedResult: false,
		},
		{
			name:           "No match REGEX_MSG_51",
			regex:          REGEX_MSG_51,
			input:          ": Create maximal test",
			expectedResult: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.regex.MatchString(tc.input)
			if actual != tc.expectedResult {
				t.Errorf("got '%t', want '%t'", actual, tc.expectedResult)
			}
		})
	}
}
