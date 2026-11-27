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

package gcp_test

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp"
	"google.golang.org/api/googleapi"
)

func TestIsNotFoundError(t *testing.T) {
	testCases := []struct {
		Name     string
		Error    error
		Expected bool
	}{
		{
			Name:     "nil error",
			Error:    nil,
			Expected: false,
		},
		{
			Name: "googleapi.Error with StatusNotFound",
			Error: &googleapi.Error{
				Code: http.StatusNotFound,
			},
			Expected: true,
		},
		{
			Name: "googleapi.Error with StatusForbidden",
			Error: &googleapi.Error{
				Code: http.StatusForbidden,
			},
			Expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			actual := gcp.IsNotFoundError(tc.Error)
			if actual != tc.Expected {
				t.Fatalf("expected '%v', got '%v'", tc.Expected, actual)
			}
		})
	}
}

func TestIsNotAuthorizedError(t *testing.T) {
	testCases := []struct {
		Name     string
		Error    error
		Expected bool
	}{
		{
			Name:     "nil error",
			Error:    nil,
			Expected: false,
		},
		{
			Name: "googleapi.Error with StatusForbidden",
			Error: &googleapi.Error{
				Code: http.StatusForbidden,
			},
			Expected: true,
		},
		{
			Name: "googleapi.Error with StatusNotFound",
			Error: &googleapi.Error{
				Code: http.StatusNotFound,
			},
			Expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			actual := gcp.IsNotAuthorizedError(tc.Error)
			if actual != tc.Expected {
				t.Fatalf("expected '%v', got '%v'", tc.Expected, actual)
			}
		})
	}
}

func TestIsServiceEndpointUnavailableError(t *testing.T) {
	testCases := []struct {
		Name         string
		EnvUniverse  string
		Error        error
		ExpectedOK   bool
		ExpectedHost string
	}{
		{
			Name:         "nil error",
			Error:        nil,
			ExpectedOK:   false,
			ExpectedHost: "",
		},
		{
			Name:         "generic error",
			Error:        fmt.Errorf("some generic failure"),
			ExpectedOK:   false,
			ExpectedHost: "",
		},
		{
			Name:         "default googleapis.com NXDOMAIN",
			Error:        fmt.Errorf("Get \"https://binaryauthorization.googleapis.com/v1/projects/sample/policy\": dial tcp: lookup binaryauthorization.googleapis.com: no such host"),
			ExpectedOK:   false,
			ExpectedHost: "",
		},
		{
			Name:         "sovereign custom universe domain NXDOMAIN string",
			Error:        fmt.Errorf("Update call failed: error fetching live state: error reading underlying resource: Get \"https://binaryauthorization.googleapis.com/v1/projects/partition:sample-project/policy?alt=json\": dial tcp: lookup binaryauthorization.custom.universe.goog: no such host"),
			ExpectedOK:   true,
			ExpectedHost: "binaryauthorization.custom.universe.goog",
		},
		{
			Name:         "lookup with resolver on port and request host",
			Error:        fmt.Errorf("Get \"https://vpcaccess.googleapis.com/v1/projects/sample/locations/us-central1/connectors\": dial tcp: lookup vpcaccess.isolated.cloud:53 on 10.0.0.2:53: no such host"),
			ExpectedOK:   true,
			ExpectedHost: "vpcaccess.isolated.cloud",
		},
		{
			Name:         "lookup with resolver on port and configured universe domain",
			EnvUniverse:  "isolated.cloud",
			Error:        fmt.Errorf("dial tcp: lookup vpcaccess.isolated.cloud:53 on 10.0.0.2:53: no such host"),
			ExpectedOK:   true,
			ExpectedHost: "vpcaccess.isolated.cloud",
		},
		{
			Name: "net.DNSError with IsNotFound true in custom domain",
			Error: &net.DNSError{
				Err:        "no such host",
				Name:       "binaryauthorization.custom.universe.goog",
				IsNotFound: true,
			},
			ExpectedOK:   true,
			ExpectedHost: "binaryauthorization.custom.universe.goog",
		},
		{
			Name: "wrapped in url.Error with custom universe domain",
			Error: &url.Error{
				Op:  "Get",
				URL: "https://binaryauthorization.custom.universe.goog/v1/foo",
				Err: &net.DNSError{
					Err:        "no such host",
					Name:       "binaryauthorization.custom.universe.goog",
					IsNotFound: true,
				},
			},
			ExpectedOK:   true,
			ExpectedHost: "binaryauthorization.custom.universe.goog",
		},
		{
			Name: "wrapped in url.Error from googleapis.com dialing custom universe",
			Error: &url.Error{
				Op:  "Get",
				URL: "https://binaryauthorization.googleapis.com/v1/foo",
				Err: &net.DNSError{
					Err:        "no such host",
					Name:       "binaryauthorization.custom.universe.goog",
					IsNotFound: true,
				},
			},
			ExpectedOK:   true,
			ExpectedHost: "binaryauthorization.custom.universe.goog",
		},
		{
			Name: "net.DNSError in googleapis.com",
			Error: &net.DNSError{
				Err:        "no such host",
				Name:       "binaryauthorization.googleapis.com",
				IsNotFound: true,
			},
			ExpectedOK:   false,
			ExpectedHost: "",
		},
		// Negative tests: metadata.google.internal
		{
			Name:         "negative: metadata.google.internal plain error",
			Error:        fmt.Errorf("dial tcp: lookup metadata.google.internal: no such host"),
			ExpectedOK:   false,
			ExpectedHost: "",
		},
		{
			Name:         "negative: metadata.google.internal with request url",
			Error:        fmt.Errorf("Get \"http://metadata.google.internal/computeMetadata/v1/instance/service-accounts/default/token\": dial tcp: lookup metadata.google.internal: no such host"),
			ExpectedOK:   false,
			ExpectedHost: "",
		},
		{
			Name: "negative: metadata.google.internal wrapped url.Error",
			Error: &url.Error{
				Op:  "Get",
				URL: "http://metadata.google.internal/computeMetadata/v1/instance",
				Err: &net.DNSError{
					Err:        "no such host",
					Name:       "metadata.google.internal",
					IsNotFound: true,
				},
			},
			ExpectedOK:   false,
			ExpectedHost: "",
		},
		// Negative tests: proxy errors
		{
			Name:         "negative: corporate proxy NXDOMAIN",
			Error:        fmt.Errorf("Get \"https://binaryauthorization.googleapis.com/v1/projects/sample/policy\": dial tcp: lookup corporate-proxy.internal:8080: no such host"),
			ExpectedOK:   false,
			ExpectedHost: "",
		},
		{
			Name:        "negative: corporate proxy NXDOMAIN with custom universe configured",
			EnvUniverse: "custom.universe.goog",
			Error:       fmt.Errorf("Get \"https://binaryauthorization.custom.universe.goog/v1/projects/sample/policy\": dial tcp: lookup corporate-proxy.internal:8080: no such host"),
			ExpectedOK:  false,
		},
		// Negative tests: endpoint overrides
		{
			Name:         "negative: custom endpoint override NXDOMAIN",
			Error:        fmt.Errorf("Get \"https://my-binauthz-mock.internal:8080/v1/projects/sample/policy\": dial tcp: lookup my-binauthz-mock.internal: no such host"),
			ExpectedOK:   false,
			ExpectedHost: "",
		},
		{
			Name:         "negative: localhost endpoint override NXDOMAIN",
			Error:        fmt.Errorf("Get \"http://localhost:8080/v1/projects/sample/policy\": dial tcp: lookup localhost: no such host"),
			ExpectedOK:   false,
			ExpectedHost: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			if tc.EnvUniverse != "" {
				t.Setenv("GOOGLE_CLOUD_UNIVERSE_DOMAIN", tc.EnvUniverse)
			} else {
				t.Setenv("GOOGLE_CLOUD_UNIVERSE_DOMAIN", "")
				t.Setenv("GOOGLE_UNIVERSE_DOMAIN", "")
			}
			host, ok := gcp.IsServiceEndpointUnavailableError(tc.Error)
			if ok != tc.ExpectedOK {
				t.Errorf("expected ok=%v, got %v for %v", tc.ExpectedOK, ok, tc.Error)
			}
			if host != tc.ExpectedHost {
				t.Errorf("expected host=%q, got %q", tc.ExpectedHost, host)
			}
		})
	}
}
