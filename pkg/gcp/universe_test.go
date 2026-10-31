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
	"io"
	"net/http"
	"strings"
	"testing"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestRewriteHostForUniverse(t *testing.T) {
	tests := []struct {
		name     string
		host     string
		path     string
		universe string
		want     string
	}{
		{
			name:     "commercial default leaves host unchanged",
			host:     "compute.googleapis.com",
			path:     "/compute/v1/projects/p1",
			universe: "googleapis.com",
			want:     "compute.googleapis.com",
		},
		{
			name:     "service subdomain rewritten to custom universe",
			host:     "storage.googleapis.com",
			path:     "/storage/v1/b/my-bucket",
			universe: "custom.universe.goog",
			want:     "storage.custom.universe.goog",
		},
		{
			name:     "www.googleapis.com compute path routed to compute.<universe>",
			host:     "www.googleapis.com",
			path:     "/compute/beta/projects/partition:sample-project/zones/u-region-1-a/instanceGroupManagers",
			universe: "custom.universe.goog",
			want:     "compute.custom.universe.goog",
		},
		{
			name:     "www.googleapis.com upload storage path routed to storage.<universe>",
			host:     "www.googleapis.com",
			path:     "/upload/storage/v1/b/my-bucket/o",
			universe: "custom.universe.goog",
			want:     "storage.custom.universe.goog",
		},
		{
			name:     "preserves port when rewriting host",
			host:     "www.googleapis.com:443",
			path:     "/compute/v1/projects/p1",
			universe: "custom.universe.goog",
			want:     "compute.custom.universe.goog:443",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := RewriteHostForUniverse(tc.host, tc.path, tc.universe); got != tc.want {
				t.Errorf("RewriteHostForUniverse(%q, %q, %q) = %q, want %q", tc.host, tc.path, tc.universe, got, tc.want)
			}
		})
	}
}

func TestUniverseDomainRoundTripper(t *testing.T) {
	var capturedHost, capturedPath, capturedQuery string
	rt := NewUniverseDomainRoundTripper(roundTripFunc(func(req *http.Request) (*http.Response, error) {
		capturedHost = req.URL.Host
		capturedPath = req.URL.Path
		capturedQuery = req.URL.RawQuery
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader("{}")),
			Header:     make(http.Header),
		}, nil
	}), "custom.universe.goog")

	req, err := http.NewRequest(http.MethodGet, "https://www.googleapis.com/compute/beta/projects/p1/zones/z1/instanceGroupManagers?alt=json", nil)
	if err != nil {
		t.Fatalf("NewRequest failed: %v", err)
	}
	if _, err := rt.RoundTrip(req); err != nil {
		t.Fatalf("RoundTrip failed: %v", err)
	}
	if capturedHost != "compute.custom.universe.goog" {
		t.Errorf("capturedHost = %q, want %q", capturedHost, "compute.custom.universe.goog")
	}
	if capturedPath != "/compute/beta/projects/p1/zones/z1/instanceGroupManagers" {
		t.Errorf("capturedPath = %q", capturedPath)
	}
	if capturedQuery != "alt=json" {
		t.Errorf("capturedQuery = %q, want %q", capturedQuery, "alt=json")
	}
}
