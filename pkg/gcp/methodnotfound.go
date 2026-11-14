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

package gcp

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"google.golang.org/api/googleapi"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// IsMethodNotFoundError returns true if the error indicates an HTTP 404 Method not found,
// which occurs when an API method or URL route is unrouted or omitted from regional proxies
// (e.g. against custom universe or sovereign endpoints).
func IsMethodNotFoundError(err error) bool {
	if err == nil {
		return false
	}

	// 1. Check for googleapi.Error (REST API responses).
	var ge *googleapi.Error
	if errors.As(err, &ge) {
		if ge.Code == http.StatusNotFound {
			// Structured signal: Check ErrorInfo reason in googleapi.Error details.
			for _, detail := range ge.Details {
				if dMap, ok := detail.(map[string]interface{}); ok {
					if reason, ok := dMap["reason"].(string); ok && strings.EqualFold(reason, "METHOD_NOT_FOUND") {
						return true
					}
				}
			}

			// Structured signal: Check ErrorItem reason in Errors.
			for _, item := range ge.Errors {
				if strings.EqualFold(item.Reason, "METHOD_NOT_FOUND") || strings.EqualFold(item.Reason, "methodNotFound") {
					return true
				}
			}

			// Structured signal: Exact method not found message from Google API reverse proxy / ESP.
			if isMethodNotFoundMessage(ge.Message) {
				return true
			}

			// Structured signal: Check JSON response body if Message was not set directly on Error.
			if ge.Body != "" {
				var bodyObj struct {
					Error struct {
						Code    int    `json:"code"`
						Message string `json:"message"`
						Status  string `json:"status"`
						Details []struct {
							Reason string `json:"reason"`
						} `json:"details"`
					} `json:"error"`
				}
				if json.Unmarshal([]byte(ge.Body), &bodyObj) == nil {
					for _, d := range bodyObj.Error.Details {
						if strings.EqualFold(d.Reason, "METHOD_NOT_FOUND") {
							return true
						}
					}
					if isMethodNotFoundMessage(bodyObj.Error.Message) {
						return true
					}
				}
			}
		}
		return false
	}

	// 2. Check for gRPC status error.
	if st, ok := status.FromError(err); ok {
		if st.Code() == codes.NotFound || st.Code() == codes.Unimplemented {
			for _, detail := range st.Details() {
				if errorInfo, ok := detail.(*errdetails.ErrorInfo); ok {
					if strings.EqualFold(errorInfo.GetReason(), "METHOD_NOT_FOUND") {
						return true
					}
				}
			}
			if isMethodNotFoundMessage(st.Message()) {
				return true
			}
		}
		return false
	}

	return false
}

func isMethodNotFoundMessage(msg string) bool {
	trimmed := strings.TrimRight(strings.TrimSpace(msg), ".")
	return strings.EqualFold(trimmed, "method not found")
}
