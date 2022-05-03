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

package execution_test

import (
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/errors"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/execution"
)

func TestRecoverWithGenericError(t *testing.T) {
	testBasicRecoverScenariosWithFunc(t, execution.RecoverWithGenericError)
}

func TestRecoverWithInternalError(t *testing.T) {
	testBasicRecoverScenariosWithFunc(t, execution.RecoverWithInternalError)
	err := testRecover(execution.RecoverWithInternalError)
	expectedType := reflect.TypeOf(&errors.InternalError{})
	actualType := reflect.TypeOf(err)
	if actualType != expectedType {
		t.Fatalf("unexpected type returned: got '%v', want' %v'", actualType, expectedType)
	}
}

func testBasicRecoverScenariosWithFunc(t *testing.T, recoverFunc func(err *error)) {
	err := testRecover(recoverFunc)
	if err == nil {
		t.Fatalf("expected an error, instead got 'nil'")
	}
	expectedRecoverSubstrings := []string{
		"panicFunc(...)",
		"midStackFunc",
		"my function is panicking!",
	}
	errMsg := err.Error()
	for _, expectedMessage := range expectedRecoverSubstrings {
		if !strings.Contains(errMsg, expectedMessage) {
			t.Errorf("expected error string to contain '%v', got: %v", expectedMessage, errMsg)
		}
	}
	unexpectedRecoverSubstrings := []string{
		getFullFunctionName(recoverFunc),
		getPackageFunctionName(recoverFunc),
		"debug.Stack",
	}
	for _, expectedMessage := range unexpectedRecoverSubstrings {
		if strings.Contains(errMsg, expectedMessage) {
			t.Errorf("expected error string to NOT contain '%v', got:\n%v", expectedMessage, errMsg)
		}
	}
}

func testRecover(recoverFunc func(err *error)) (err error) {
	defer recoverFunc(&err)
	midStackFunc()
	return err
}

func midStackFunc() {
	panicFunc()
}

func panicFunc() {
	panic("my function is panicking!")
}

// getPackageFunctionName returns the package level function name, ex: execution.RecoverWithInternalError
func getPackageFunctionName(i interface{}) string {
	fullName := getFullFunctionName(i)
	return fullName[strings.LastIndex(fullName, "/")+1:]
}

// getFullFunctionName returns the fully qualified function name, ex:
//   github.com/GoogleCloudPlatform/k8s-config-connector/pkg/execution.RecoverWithInternalError
func getFullFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
