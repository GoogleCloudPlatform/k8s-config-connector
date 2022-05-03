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

package valutil_test

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/valutil"
)

func TestIsDefaultValue(t *testing.T) {
	testCases := []struct {
		Name           string
		Value          interface{}
		ExpectedResult bool
	}{
		{
			Name:           "Default string",
			Value:          "",
			ExpectedResult: true,
		},
		{
			Name:           "String value",
			Value:          "a",
			ExpectedResult: false,
		},
		{
			Name:           "Default bool",
			Value:          false,
			ExpectedResult: true,
		},
		{
			Name:           "Bool value",
			Value:          true,
			ExpectedResult: false,
		},
		{
			Name:           "Pointer to string",
			Value:          new(string),
			ExpectedResult: true,
		},
		{
			Name:           "Pointer to bool",
			Value:          new(bool),
			ExpectedResult: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := valutil.IsDefaultValue(tc.Value)
			if result != tc.ExpectedResult {
				t.Fatalf("unexpected result: got '%v', want '%v'", result, tc.ExpectedResult)
			}
		})
	}
}
