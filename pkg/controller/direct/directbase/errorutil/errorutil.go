// Copyright 2026 Google LLC
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

package errorutil

import (
	"errors"
	"net/http"
	"strings"

	"github.com/googleapis/gax-go/v2/apierror"
	"google.golang.org/api/googleapi"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var transientReasons = map[string]bool{
	"RESOURCE_IN_USE_BY_ANOTHER_RESOURCE": true,
	"RESOURCE_IN_USE_BY_OTHER_RESOURCE":   true,
	"RESOURCE_NOT_READY":                  true,
	"RESOURCE_BEING_MODIFIED":             true,
	"RESOURCE_IN_USE":                     true,
	"RESOURCE_USAGE_LIMIT_REACHED":        true,
	"ZONE_RESOURCE_POOL_EXHAUSTED":        true,
	"OPERATION_IN_PROGRESS":               true,
	"RESOURCE_ALREADY_EXISTS":             true,
	"RATE_LIMIT_EXCEEDED":                 true,
	"SERVICE_UNAVAILABLE":                 true,
	"DEADLINE_EXCEEDED":                   true,
	"LOCKED":                              true,
	"CONCURRENT_MUTATION":                 true,
	"ABORTED":                             true,
}

var transientSubstrings = []string{
	"resource_in_use_by_another_resource",
	"resource_in_use_by_other_resource",
	"resource in use",
	"in use by another resource",
	"in use by other resource",
	"in use by",
	"resource is not ready",
	"resource not ready",
	"operation in progress",
	"operation is already in progress",
	"is currently being modified",
	"dependency lock",
	"try again",
	"temporarily unavailable",
	"temporary failure",
	"transient",
	"concurrent mutation",
	"quota exceeded",
	"rate limit",
}

var validArgumentReasons = map[string]bool{
	"INVALID_ARGUMENT":       true,
	"FIELD_VIOLATION":        true,
	"BAD_REQUEST":            true,
	"IMMUTABLE_FIELD":        true,
	"VALUE_OUT_OF_RANGE":     true,
	"OUT_OF_BOUNDS":          true,
	"INVALID_PARAMETER":      true,
	"SCHEMA_VIOLATION":       true,
	"INVALID_FIELD_VALUE":    true,
	"MISSING_REQUIRED_FIELD": true,
	"MALFORMED_NAME":         true,
	"NAME_VIOLATION":         true,
	"MUTATION_NOT_ALLOWED":   true,
	"CANNOT_UPDATE_FIELD":    true,
	"invalid":                true,
	"badRequest":             true,
	"invalidParameter":       true,
	"required":               true,
	"unknownField":           true,
	"badContent":             true,
	"parseError":             true,
	"fieldViolation":         true,
	"conditionNotMet":        true,
	"outOfBounds":            true,
}

var invalidArgumentMessageSubstrings = []string{
	"invalid",
	"bad request",
	"field is required",
	"missing required",
	"required field",
	"immutable",
	"cannot be updated",
	"cannot be changed",
	"cannot be modified",
	"read-only",
	"must match",
	"regex",
	"regular expression",
	"out of range",
	"must be between",
	"exceeds maximum",
	"less than minimum",
	"bounds",
	"unknown field",
	"unrecognized field",
	"unexpected field",
	"unsupported",
	"malformed",
	"syntax error",
	"parse error",
	"cannot specify both",
	"mutually exclusive",
}

// ExtractGRPCStatus returns the gRPC Status extracted from err if present.
func ExtractGRPCStatus(err error) (*status.Status, bool) {
	if err == nil {
		return nil, false
	}
	if s, ok := status.FromError(err); ok {
		return s, true
	}
	var apiErr *apierror.APIError
	if errors.As(err, &apiErr) {
		if s := apiErr.GRPCStatus(); s != nil {
			return s, true
		}
	}
	for curr := err; curr != nil; curr = errors.Unwrap(curr) {
		if s, ok := status.FromError(curr); ok {
			return s, true
		}
	}
	return nil, false
}

// ExtractGoogleAPIError returns *googleapi.Error if present in the error chain.
func ExtractGoogleAPIError(err error) (*googleapi.Error, bool) {
	if err == nil {
		return nil, false
	}
	var gerr *googleapi.Error
	if errors.As(err, &gerr) {
		return gerr, true
	}
	return nil, false
}

// ExtractAPIError returns *apierror.APIError if present in the error chain.
func ExtractAPIError(err error) (*apierror.APIError, bool) {
	if err == nil {
		return nil, false
	}
	var apiErr *apierror.APIError
	if errors.As(err, &apiErr) {
		return apiErr, true
	}
	return nil, false
}

// ExtractHTTPCode returns the HTTP status code if the error represents an HTTP response error.
func ExtractHTTPCode(err error) (int, bool) {
	if err == nil {
		return 0, false
	}
	if gerr, ok := ExtractGoogleAPIError(err); ok {
		return gerr.Code, true
	}
	if apiErr, ok := ExtractAPIError(err); ok {
		if code := apiErr.HTTPCode(); code != 0 {
			return code, true
		}
	}
	return 0, false
}

// ExtractGRPCStatusCode returns the gRPC Code if the error represents a gRPC status error.
func ExtractGRPCStatusCode(err error) (codes.Code, bool) {
	if err == nil {
		return codes.OK, false
	}
	if s, ok := ExtractGRPCStatus(err); ok {
		return s.Code(), true
	}
	return codes.OK, false
}

// IsTerminalError checks if the error represents a permanent client error
// (such as invalid arguments, malformed syntax, or attempts to mutate immutable fields)
// that should halt workqueue retries under terminal-error-mode: "builtin".
func IsTerminalError(err error) bool {
	return IsTerminalInvalidArgument(err)
}

// IsTerminalInvalidArgument verifies whether the error is a genuine InvalidArgument / HTTP 400 client error,
// explicitly filtering out transient dependency locks, rate limits, or 5xx server errors.
func IsTerminalInvalidArgument(err error) bool {
	if err == nil {
		return false
	}

	st, hasGRPC := ExtractGRPCStatus(err)
	gerr, hasGoogleAPI := ExtractGoogleAPIError(err)
	apiErr, hasAPIError := ExtractAPIError(err)

	isInvalidArgument := false

	if hasGRPC && st.Code() == codes.InvalidArgument {
		isInvalidArgument = true
	}
	if hasGoogleAPI && gerr.Code == http.StatusBadRequest {
		isInvalidArgument = true
	}
	if hasAPIError {
		if apiErr.HTTPCode() == http.StatusBadRequest {
			isInvalidArgument = true
		}
		if apiErr.GRPCStatus() != nil && apiErr.GRPCStatus().Code() == codes.InvalidArgument {
			isInvalidArgument = true
		}
	}

	// Must be HTTP 400 or gRPC InvalidArgument (3).
	if !isInvalidArgument {
		return false
	}

	// Filter out transient dependency locks and retryable states from ErrorInfo / details
	if hasGRPC && st != nil {
		for _, detail := range st.Details() {
			if errorInfo, ok := detail.(*errdetails.ErrorInfo); ok {
				if transientReasons[errorInfo.GetReason()] {
					return false
				}
			}
		}
	}
	if hasAPIError && apiErr != nil {
		if errorInfo := apiErr.Details().ErrorInfo; errorInfo != nil {
			if transientReasons[errorInfo.GetReason()] {
				return false
			}
		}
	}
	if hasGoogleAPI && gerr != nil {
		for _, item := range gerr.Errors {
			if transientReasons[item.Reason] {
				return false
			}
			msgLower := strings.ToLower(item.Message)
			for _, sub := range transientSubstrings {
				if strings.Contains(msgLower, sub) {
					return false
				}
			}
		}
		bodyLower := strings.ToLower(gerr.Body)
		for _, sub := range transientSubstrings {
			if strings.Contains(bodyLower, sub) {
				return false
			}
		}
	}

	errStrLower := strings.ToLower(err.Error())
	for _, sub := range transientSubstrings {
		if strings.Contains(errStrLower, sub) {
			return false
		}
	}

	// Verify genuine invalid arguments
	// 1. Check for errdetails.BadRequest (structured field violations)
	if hasGRPC && st != nil {
		for _, detail := range st.Details() {
			if _, ok := detail.(*errdetails.BadRequest); ok {
				return true
			}
			if errorInfo, ok := detail.(*errdetails.ErrorInfo); ok {
				if validArgumentReasons[errorInfo.GetReason()] {
					return true
				}
			}
		}
	}
	if hasAPIError && apiErr != nil {
		if apiErr.Details().BadRequest != nil {
			return true
		}
		if errorInfo := apiErr.Details().ErrorInfo; errorInfo != nil {
			if validArgumentReasons[errorInfo.GetReason()] {
				return true
			}
		}
	}

	// 2. Check googleapi.Error items
	if hasGoogleAPI && gerr != nil {
		for _, item := range gerr.Errors {
			if validArgumentReasons[item.Reason] {
				return true
			}
		}
	}

	// 3. Check for invalid argument message indicators or canonical InvalidArgument gRPC code
	if hasGRPC && st != nil && st.Code() == codes.InvalidArgument {
		return true
	}

	for _, sub := range invalidArgumentMessageSubstrings {
		if strings.Contains(errStrLower, sub) {
			return true
		}
	}

	return false
}
