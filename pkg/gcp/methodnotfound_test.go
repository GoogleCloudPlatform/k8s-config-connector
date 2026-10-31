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

package gcp_test

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp"
	"google.golang.org/api/googleapi"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestIsMethodNotFoundError(t *testing.T) {
	// Simulated real response body from sovereign endpoint reverse proxy / ESP
	sovereignResponseBody := `{"error": {"code": 404, "message": "Method not found.", "status": "NOT_FOUND"}}`

	grpcStatusWithReason, err := status.New(codes.NotFound, "Resource error").WithDetails(&errdetails.ErrorInfo{
		Reason: "METHOD_NOT_FOUND",
		Domain: "googleapis.com",
	})
	if err != nil {
		t.Fatalf("failed to create gRPC status with details: %v", err)
	}

	testCases := []struct {
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
			name: "sovereign endpoint REST response with Body",
			err: &googleapi.Error{
				Code:    http.StatusNotFound,
				Message: "Method not found.",
				Body:    sovereignResponseBody,
			},
			expected: true,
		},
		{
			name: "googleapi 404 Method not found with trailing period",
			err: &googleapi.Error{
				Code:    http.StatusNotFound,
				Message: "Method not found.",
			},
			expected: true,
		},
		{
			name: "googleapi 404 Method not found without trailing period",
			err: &googleapi.Error{
				Code:    http.StatusNotFound,
				Message: "Method not found",
			},
			expected: true,
		},
		{
			name: "wrapped googleapi 404 Method not found",
			err: fmt.Errorf("error applying desired state: %w", &googleapi.Error{
				Code:    http.StatusNotFound,
				Message: "Method not found.",
			}),
			expected: true,
		},
		{
			name: "googleapi 404 with structured ErrorInfo details",
			err: &googleapi.Error{
				Code:    http.StatusNotFound,
				Message: "Endpoint not routed",
				Details: []interface{}{
					map[string]interface{}{
						"@type":  "type.googleapis.com/google.rpc.ErrorInfo",
						"reason": "METHOD_NOT_FOUND",
						"domain": "googleapis.com",
					},
				},
			},
			expected: true,
		},
		{
			name: "googleapi 404 with structured ErrorItem reason",
			err: &googleapi.Error{
				Code: http.StatusNotFound,
				Errors: []googleapi.ErrorItem{
					{
						Reason:  "METHOD_NOT_FOUND",
						Message: "The requested method is not available in this region.",
					},
				},
			},
			expected: true,
		},
		{
			name:     "gRPC status with codes.NotFound and message Method not found.",
			err:      status.Error(codes.NotFound, "Method not found."),
			expected: true,
		},
		{
			name:     "gRPC status with codes.NotFound and structured ErrorInfo reason",
			err:      grpcStatusWithReason.Err(),
			expected: true,
		},
		{
			name: "regular resource 404 not found should not match",
			err: &googleapi.Error{
				Code:    http.StatusNotFound,
				Message: "Resource 'projects/sample/locations/u-region-1/serverTlsPolicies/test' was not found.",
			},
			expected: false,
		},
		{
			name: "404 mentioning method in context of a resource should not match",
			err: &googleapi.Error{
				Code:    http.StatusNotFound,
				Message: "Payment method not found for account.",
				Errors: []googleapi.ErrorItem{
					{
						Reason:  "notFound",
						Message: "Payment method not found for account.",
					},
				},
			},
			expected: false,
		},
		{
			name: "400 with Method not found should not match",
			err: &googleapi.Error{
				Code:    http.StatusBadRequest,
				Message: "Method not found.",
			},
			expected: false,
		},
		{
			name: "500 Internal Server Error should not match",
			err: &googleapi.Error{
				Code:    http.StatusInternalServerError,
				Message: "Internal server error.",
			},
			expected: false,
		},
		{
			name:     "generic string error without structured signal should not match",
			err:      errors.New("googleapi: Error 404: Method not found."),
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := gcp.IsMethodNotFoundError(tc.err); got != tc.expected {
				t.Errorf("IsMethodNotFoundError(%v) = %v, want %v", tc.err, got, tc.expected)
			}
		})
	}
}
