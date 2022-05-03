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

package gcp_test

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp"
)

func TestFullResourceNameToShortName(t *testing.T) {
	testCases := []struct {
		Name           string
		Input          string
		ExpectedResult string
	}{
		{"short name", "my-name", "my-name"},
		{"pubsub topic", "projects/my-project-id/topics/my-topic-name", "my-topic-name"},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := gcp.FullResourceNameToShortName(tc.Input)
			if result != tc.ExpectedResult {
				t.Errorf("unexpected result, got '%v', want '%v'", result, tc.ExpectedResult)
			}
		})
	}
}
