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

package codebot

import (
	"reflect"
	"sort"
	"testing"
)

func TestGetFilteredTools(t *testing.T) {
	testCases := []struct {
		Name           string
		IncludeFlag    string
		ExcludeFlag    string
		ExpectedResult []string
		ExpectedErr    bool
	}{
		{
			Name:           "Include all tools",
			IncludeFlag:    "",
			ExcludeFlag:    "",
			ExpectedResult: []string{"CreateFile", "EditFile", "FindInWorkspace", "ListInWorkspace", "ReadFile", "RunTerminalCommand", "RunShellCommand", "VerifyCode"},
		},

		{
			Name:           "Include by regex",
			IncludeFlag:    "File",
			ExcludeFlag:    "",
			ExpectedResult: []string{"CreateFile", "EditFile", "ReadFile"},
		},

		{
			Name:           "exclude by regex 1",
			IncludeFlag:    "",
			ExcludeFlag:    "FindInWorkspace|ListInWorkspace",
			ExpectedResult: []string{"CreateFile", "EditFile", "ReadFile", "RunTerminalCommand", "RunShellCommand", "VerifyCode"},
		},
		{
			Name:           "exclude by regex 2",
			IncludeFlag:    "",
			ExcludeFlag:    "Workspace",
			ExpectedResult: []string{"CreateFile", "EditFile", "ReadFile", "RunTerminalCommand", "RunShellCommand", "VerifyCode"},
		},

		{
			Name:           "include and exclude",
			IncludeFlag:    "File",
			ExcludeFlag:    "ReadFile",
			ExpectedResult: []string{"CreateFile", "EditFile"}, // ReadFile excluded.
		},

		{
			Name:        "invalid regex, no panic",
			IncludeFlag: "*File", // Invalid regex, should not panic.
			ExpectedErr: true,    // Expect an error due to invalid regex.
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			filter, err := NewCustomRegexFilter(tc.IncludeFlag, tc.ExcludeFlag)
			if err != nil {
				if tc.ExpectedErr == false {
					t.Fatalf("NewCustomRegexFilter failed: IncludeFlag %v, ExcludeFlag %v, %v", tc.IncludeFlag, tc.ExcludeFlag, err)
				} else {
					return
				}
			}
			tools := GetFilteredTools(filter)
			actualToolNames := []string{}
			for _, tool := range tools {
				actualToolNames = append(actualToolNames, tool.BuildFunctionDefinition().Name)
			}

			sort.Slice(actualToolNames, func(i, j int) bool {
				return actualToolNames[i] < actualToolNames[j]
			})
			sort.Slice(tc.ExpectedResult, func(i, j int) bool {
				return tc.ExpectedResult[i] < tc.ExpectedResult[j]
			})

			if !reflect.DeepEqual(actualToolNames, tc.ExpectedResult) {
				t.Fatalf("CustomRegexFilter mismatch: got '%v', want '%v'", actualToolNames, tc.ExpectedResult)
			}
		})
	}
}
