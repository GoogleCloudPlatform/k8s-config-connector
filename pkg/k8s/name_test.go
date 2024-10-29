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

package k8s_test

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
)

func TestValueToDNSSubdomainName(t *testing.T) {
	testCases := []struct {
		Name           string
		Value          string
		ExpectedResult string
	}{
		{
			Name:           "empty string",
			Value:          "",
			ExpectedResult: "a",
		},
		{
			Name:           "Strip non-alphanumeric",
			Value:          "this-name-contains-#chars;_that#should-be-replaced",
			ExpectedResult: "this-name-contains--chars--that-should-be-replaced",
		},
		{
			Name:           "Periods and hyphens",
			Value:          "this.is-legal",
			ExpectedResult: "this.is-legal",
		},
		{
			Name:           "Number",
			Value:          "0123456789",
			ExpectedResult: "0123456789",
		},
		{
			Name:           "hyphen prefix",
			Value:          "-value",
			ExpectedResult: "a-value",
		},
		{
			Name:           "hyphen prefix period suffix",
			Value:          "-value.",
			ExpectedResult: "a-value.a",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := k8s.ValueToDNSSubdomainName(tc.Value)
			if result != tc.ExpectedResult {
				t.Fatalf("result mismatch: got '%v', want '%v'", result, tc.ExpectedResult)
			}
		})
	}
}
