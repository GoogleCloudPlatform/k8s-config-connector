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
	"fmt"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp"
	"google.golang.org/api/googleapi"
)

func TestIsNotFoundError(t *testing.T) {
	testCases := []struct {
		Name           string
		Error          error
		ExpectedResult bool
	}{
		{"GCP NotFound", &googleapi.Error{Code: 404}, true},
		{"GCP NotAuthorized", &googleapi.Error{Code: 403}, false},
		{"Generic", fmt.Errorf("my generic error"), false},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := gcp.IsNotFoundError(tc.Error)
			if result != tc.ExpectedResult {
				t.Errorf("unexpected result for gcp.IsNotFoundError('%v'): got '%v', want '%v'", tc.Error, result, tc.ExpectedResult)
			}
		})
	}
}

func TestIsNotAuthorizedError(t *testing.T) {
	testCases := []struct {
		Name           string
		Error          error
		ExpectedResult bool
	}{
		{"GCP NotAuthorized", &googleapi.Error{Code: 403}, true},
		{"GCP NotFound", &googleapi.Error{Code: 404}, false},
		{"Generic", fmt.Errorf("my generic error"), false},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := gcp.IsNotAuthorizedError(tc.Error)
			if result != tc.ExpectedResult {
				t.Errorf("unexpected result for gcp.IsNotAuthorizedError('%v'): got '%v', want '%v'", tc.Error, result, tc.ExpectedResult)
			}
		})
	}
}
