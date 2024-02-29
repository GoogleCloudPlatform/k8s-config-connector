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

package gcp

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"google.golang.org/api/googleapi"
	sqladmin "google.golang.org/api/sqladmin/v1beta4"
)

func IsNotAuthorizedError(err error) bool {
	return isGoogleErrorWithCode(err, http.StatusForbidden)
}

func IsNotFoundError(err error) bool {
	return isGoogleErrorWithCode(err, http.StatusNotFound)
}

func isGoogleErrorWithCode(err error, code int) bool {
	if err == nil {
		return false
	}
	ge := &googleapi.Error{}
	if errors.As(err, &ge) {
		return ge.Code == code
	}

	return false
}

func FormatSQLOperationErrorMessages(opErrs []*sqladmin.OperationError) string {
	var errMessages []string
	for _, opErr := range opErrs {
		errMessages = append(errMessages, fmt.Sprintf("%v %v", opErr.Code, opErr.Message))
	}
	return strings.Join(errMessages, ", ")
}
