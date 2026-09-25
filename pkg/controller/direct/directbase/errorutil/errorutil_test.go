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
	"fmt"
	"net/http"
	"testing"

	"github.com/googleapis/gax-go/v2/apierror"
	"google.golang.org/api/googleapi"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestIsTerminalError(t *testing.T) {
	// 1. gRPC status with BadRequest FieldViolations
	stWithBadRequest, err := status.New(codes.InvalidArgument, "Invalid request payload").WithDetails(
		&errdetails.BadRequest{
			FieldViolations: []*errdetails.BadRequest_FieldViolation{
				{
					Field:       "spec.name",
					Description: "Field must match regular expression ^[a-z0-9-]+$",
				},
			},
		},
	)
	if err != nil {
		t.Fatalf("failed to create status with details: %v", err)
	}

	// 2. gRPC status with ErrorInfo indicating lock / transient
	stWithLockErrorInfo, err := status.New(codes.InvalidArgument, "Resource in use").WithDetails(
		&errdetails.ErrorInfo{
			Reason: "RESOURCE_IN_USE_BY_ANOTHER_RESOURCE",
			Domain: "compute.googleapis.com",
		},
	)
	if err != nil {
		t.Fatalf("failed to create status with details: %v", err)
	}

	// 3. gRPC status with ErrorInfo indicating client error
	stWithClientErrorInfo, err := status.New(codes.InvalidArgument, "Bad request").WithDetails(
		&errdetails.ErrorInfo{
			Reason: "FIELD_VIOLATION",
			Domain: "googleapis.com",
		},
	)
	if err != nil {
		t.Fatalf("failed to create status with details: %v", err)
	}

	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "nil error",
			err:      nil,
			expected: false,
		},
		{
			name:     "generic go error",
			err:      errors.New("unexpected internal error"),
			expected: false,
		},
		{
			name:     "gRPC InvalidArgument with BadRequest field violations",
			err:      stWithBadRequest.Err(),
			expected: true,
		},
		{
			name:     "wrapped gRPC InvalidArgument with BadRequest field violations",
			err:      fmt.Errorf("error creating resource: %w", stWithBadRequest.Err()),
			expected: true,
		},
		{
			name:     "gRPC InvalidArgument plain status error",
			err:      status.Error(codes.InvalidArgument, "name 'INVALID_NAME' must match regex ^[a-z0-9-]+$"),
			expected: true,
		},
		{
			name:     "gRPC InvalidArgument with transient lock ErrorInfo",
			err:      stWithLockErrorInfo.Err(),
			expected: false,
		},
		{
			name:     "gRPC InvalidArgument with client error ErrorInfo",
			err:      stWithClientErrorInfo.Err(),
			expected: true,
		},
		{
			name:     "gRPC Internal error (500)",
			err:      status.Error(codes.Internal, "internal server error"),
			expected: false,
		},
		{
			name:     "gRPC Unavailable error (503)",
			err:      status.Error(codes.Unavailable, "service unavailable"),
			expected: false,
		},
		{
			name:     "gRPC PermissionDenied error (403)",
			err:      status.Error(codes.PermissionDenied, "IAM permission denied"),
			expected: false,
		},
		{
			name:     "gRPC NotFound error (404)",
			err:      status.Error(codes.NotFound, "resource not found"),
			expected: false,
		},
		{
			name:     "gRPC ResourceExhausted (429)",
			err:      status.Error(codes.ResourceExhausted, "quota exceeded"),
			expected: false,
		},
		{
			name: "googleapi.Error 400 with invalidParameter reason",
			err: &googleapi.Error{
				Code:    http.StatusBadRequest,
				Message: "Invalid parameter value",
				Errors: []googleapi.ErrorItem{
					{
						Reason:  "invalidParameter",
						Message: "Invalid value for 'retention_days': must be between 1 and 365",
					},
				},
			},
			expected: true,
		},
		{
			name: "googleapi.Error 400 with badRequest reason",
			err: &googleapi.Error{
				Code:    http.StatusBadRequest,
				Message: "Bad Request: field 'tier' is required",
				Errors: []googleapi.ErrorItem{
					{
						Reason:  "badRequest",
						Message: "Field 'tier' is required",
					},
				},
			},
			expected: true,
		},
		{
			name: "googleapi.Error 400 with transient RESOURCE_IN_USE_BY_ANOTHER_RESOURCE reason",
			err: &googleapi.Error{
				Code:    http.StatusBadRequest,
				Message: "The resource is currently in use",
				Errors: []googleapi.ErrorItem{
					{
						Reason:  "RESOURCE_IN_USE_BY_ANOTHER_RESOURCE",
						Message: "Resource is in use by another resource",
					},
				},
			},
			expected: false,
		},
		{
			name: "googleapi.Error 400 with transient message substring",
			err: &googleapi.Error{
				Code:    http.StatusBadRequest,
				Message: "Resource is not ready, try again later",
			},
			expected: false,
		},
		{
			name: "googleapi.Error 500",
			err: &googleapi.Error{
				Code:    http.StatusInternalServerError,
				Message: "Internal backend failure",
			},
			expected: false,
		},
		{
			name: "googleapi.Error 403 (IAM / Forbidden)",
			err: &googleapi.Error{
				Code:    http.StatusForbidden,
				Message: "Permission denied on resource",
			},
			expected: false,
		},
		{
			name: "googleapi.Error 409 (Conflict)",
			err: &googleapi.Error{
				Code:    http.StatusConflict,
				Message: "Resource already exists",
			},
			expected: false,
		},
		{
			name: "gax apierror with 400 and BadRequest details",
			err: func() error {
				apiErr, _ := apierror.ParseError(stWithBadRequest.Err(), false)
				return apiErr
			}(),
			expected: true,
		},
		{
			name: "gax apierror with 400 and Lock details",
			err: func() error {
				apiErr, _ := apierror.ParseError(stWithLockErrorInfo.Err(), false)
				return apiErr
			}(),
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsTerminalError(tc.err)
			if got != tc.expected {
				t.Errorf("IsTerminalError(%v) = %v, want %v", tc.err, got, tc.expected)
			}
		})
	}
}

func TestExtractCodes(t *testing.T) {
	st := status.New(codes.InvalidArgument, "bad argument")
	gerr := &googleapi.Error{Code: http.StatusBadRequest, Message: "bad request"}

	if code, ok := ExtractGRPCStatusCode(st.Err()); !ok || code != codes.InvalidArgument {
		t.Errorf("ExtractGRPCStatusCode() = (%v, %v), want (%v, true)", code, ok, codes.InvalidArgument)
	}

	if code, ok := ExtractHTTPCode(gerr); !ok || code != http.StatusBadRequest {
		t.Errorf("ExtractHTTPCode() = (%v, %v), want (%v, true)", code, ok, http.StatusBadRequest)
	}

	wrapped := fmt.Errorf("wrap: %w", gerr)
	if code, ok := ExtractHTTPCode(wrapped); !ok || code != http.StatusBadRequest {
		t.Errorf("ExtractHTTPCode(wrapped) = (%v, %v), want (%v, true)", code, ok, http.StatusBadRequest)
	}
}
