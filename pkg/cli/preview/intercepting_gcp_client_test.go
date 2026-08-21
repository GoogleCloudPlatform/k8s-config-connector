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
