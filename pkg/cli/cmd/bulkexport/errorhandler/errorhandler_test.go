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

package errorhandler

import (
	"reflect"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/bulkexport/parameters"
)

func TestNewErrorHandler(t *testing.T) {
	testCases := []struct {
		Name         string
		OnErrorParam string
		ExpectedType reflect.Type
	}{
		{
			Name:         "Halt",
			OnErrorParam: "halt",
			ExpectedType: reflect.TypeOf(&haltHandler{}),
		},
		{
			Name:         "Continue",
			OnErrorParam: "continue",
			ExpectedType: reflect.TypeOf(&continueHandler{}),
		},
		{
			Name:         "Ignore",
			OnErrorParam: "ignore",
			ExpectedType: reflect.TypeOf(&ignoreHandler{}),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			params := parameters.Parameters{
				OnError: tc.OnErrorParam,
			}
			handler, err := NewErrorHandler(&params)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if reflect.TypeOf(handler) != tc.ExpectedType {
				t.Fatalf("type mismatch: got '%v', want '%v'", reflect.TypeOf(handler).Name(), tc.ExpectedType.Name())
			}
		})
	}
}

func TestNewErrorHandlerInvalidParam(t *testing.T) {
	params := parameters.Parameters{
		OnError: "my-invalid-error-parameter",
	}
	handler, err := NewErrorHandler(&params)
	if err == nil {
		t.Fatalf("value for err: got 'nil', want 'an error'")
	}
	expectedErrorMsg := "unknown 'on-error' option 'my-invalid-error-parameter', must be one of: continue, halt, ignore"
	if err.Error() != expectedErrorMsg {
		t.Fatalf("unexpected error message: got '%v', want '%v'", err.Error(), expectedErrorMsg)
	}
	if handler != nil {
		t.Fatalf("unexpected non-nil value for handler")
	}
}
