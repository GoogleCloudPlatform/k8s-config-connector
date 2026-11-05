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

package preview

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestCheckGCPRequestIsAllowed(t *testing.T) {
	client := &interceptingGCPClient{}

	tests := []struct {
		name     string
		method   string
		path     string
		rawQuery string
		allowed  bool
	}{
		{
			name:    "GET request allowed",
			method:  "GET",
			path:    "/v1/projects/foo",
			allowed: true,
		},
		{
			name:    "POST request blocked by default",
			method:  "POST",
			path:    "/v1/projects/foo",
			allowed: false,
		},
		{
			name:    "POST getIamPolicy allowed",
			method:  "POST",
			path:    "/v1/projects/project-id/serviceAccounts/user@project-id.iam.gserviceaccount.com:getIamPolicy",
			allowed: true,
		},
		{
			name:     "POST getIamPolicy with query params allowed",
			method:   "POST",
			path:     "/v1/projects/foo:getIamPolicy",
			rawQuery: "alt=json&prettyPrint=false",
			allowed:  true,
		},
		{
			name:    "POST getOrgPolicy allowed",
			method:  "POST",
			path:    "/v1/projects/foo:getOrgPolicy",
			allowed: true,
		},
		{
			name:    "POST other custom method blocked",
			method:  "POST",
			path:    "/v1/projects/foo:setIamPolicy",
			allowed: false,
		},
		{
			name:    "POST with multiple colons blocked",
			method:  "POST",
			path:    "/v1/projects/foo:bar:getIamPolicy",
			allowed: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := &http.Request{
				Method: tc.method,
				URL: &url.URL{
					Path:     tc.path,
					RawQuery: tc.rawQuery,
				},
			}
			got := client.checkGCPRequestIsAllowed(req)
			if got != tc.allowed {
				t.Errorf("checkGCPRequestIsAllowed(%s %s) = %v, want %v", tc.method, tc.path, got, tc.allowed)
			}
		})
	}
}

func TestExtractBlockedGCPError(t *testing.T) {
	origErr := BlockedGCPError{
		Method: "POST",
		URL:    "https://example.com/v1/resource",
	}

	tests := []struct {
		name       string
		err        error
		wantOk     bool
		wantMethod string
	}{
		{
			name:   "nil error",
			err:    nil,
			wantOk: false,
		},
		{
			name:       "direct value error",
			err:        origErr,
			wantOk:     true,
			wantMethod: "POST",
		},
		{
			name:       "direct pointer error",
			err:        &origErr,
			wantOk:     true,
			wantMethod: "POST",
		},
		{
			name:       "wrapped value error",
			err:        fmt.Errorf("Update call failed: %w", origErr),
			wantOk:     true,
			wantMethod: "POST",
		},
		{
			name:       "wrapped pointer error",
			err:        fmt.Errorf("Update call failed: %w", &origErr),
			wantOk:     true,
			wantMethod: "POST",
		},
		{
			name:       "string wrapped json format (terraform style)",
			err:        fmt.Errorf("TF provider error: %s", origErr.Error()),
			wantOk:     true,
			wantMethod: "POST",
		},
		{
			name:   "regular non-blocked error",
			err:    fmt.Errorf("mapping error: invalid field"),
			wantOk: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := ExtractBlockedGCPError(tc.err)
			if ok != tc.wantOk {
				t.Errorf("ExtractBlockedGCPError() ok = %v, want %v", ok, tc.wantOk)
			}
			if ok && got.Method != tc.wantMethod {
				t.Errorf("ExtractBlockedGCPError() Method = %v, want %v", got.Method, tc.wantMethod)
			}
		})
	}
}

func TestIsBlockedError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "nil error",
			err:  nil,
			want: false,
		},
		{
			name: "direct BlockedGCPError",
			err:  BlockedGCPError{Method: "PUT", URL: "https://foo"},
			want: true,
		},
		{
			name: "wrapped BlockedGCPError",
			err:  fmt.Errorf("Update call failed: %w", BlockedGCPError{Method: "PUT", URL: "https://foo"}),
			want: true,
		},
		{
			name: "kube blocked error",
			err:  fmt.Errorf("\"update\" blocked in preview mode"),
			want: true,
		},
		{
			name: "kube status blocked error wrapped",
			err:  fmt.Errorf("error updating status: %w", fmt.Errorf("\"status.update\" blocked in preview mode")),
			want: true,
		},
		{
			name: "grpc blocked error",
			err:  fmt.Errorf("GRPC method blocked by InterceptingGCPClient"),
			want: true,
		},
		{
			name: "call to GCP blocked string",
			err:  fmt.Errorf("something failed: call to GCP blocked (method=POST, url=http://foo)"),
			want: true,
		},
		{
			name: "regular reconcile error",
			err:  fmt.Errorf("mapping error: missing required field"),
			want: false,
		},
		{
			name: "reference resolution error",
			err:  fmt.Errorf("reference not found: bucket foo"),
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsBlockedError(tc.err)
			if got != tc.want {
				t.Errorf("IsBlockedError(%v) = %v, want %v", tc.err, got, tc.want)
			}
		})
	}
}
