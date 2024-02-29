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

package execution

import (
	"errors"
	"fmt"
	"runtime/debug"
	"strings"

	kccerrors "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/errors"
)

// RecoverWithGenericError is a general purpose function for recovering from panics. A useful error is written to 'err'. See
// RecoverWithInternalError for recovering inside a controller Reconcile loop
func RecoverWithGenericError(err *error) {
	if rec := recover(); rec != nil {
		handleRecovery(err, rec, errors.New)
	}
}

// Recovers with the 'err' value being filled in with an InternalError, used by controllers to enable metric collection
func RecoverWithInternalError(err *error) {
	newError := func(message string) error {
		return kccerrors.NewInternalError("panic", message)
	}
	if rec := recover(); rec != nil {
		handleRecovery(err, rec, newError)
	}
}

type NewErrorFunc func(message string) error

func handleRecovery(err *error, rec interface{}, newError NewErrorFunc) {
	stack := debug.Stack()
	formattedStack := string(stack)
	// remove the three functions from the stack trace:
	//   - debug.Stack(...)
	//   - execution.handleRecovery()
	//   - calling function (i.e. one of the Recover functions in this package)
	// each function gets two lines of output
	for i := 0; i < 6; i++ {
		formattedStack = safeRemoveFirstLine(formattedStack)
	}
	// indent all lines of the stack trace to make it easier to follow in a log
	formattedStack = strings.ReplaceAll(formattedStack, "\n", "\n\t")
	*err = newError(fmt.Sprintf("observed a panic: %+v\n%v", rec, formattedStack))
}

func safeRemoveFirstLine(value string) string {
	idx := strings.Index(value, "\n")
	if idx >= 0 {
		if len(value) > idx+1 {
			idx++
		}
		return value[idx:]
	}
	return value
}
