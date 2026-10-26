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

package discovery

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestServiceNameForCRD(t *testing.T) {
	tests := []struct {
		crd  string
		want string
	}{
		{"storagebuckets.storage.cnrm.cloud.google.com", "storage.googleapis.com"},
		{"sqlinstances.sql.cnrm.cloud.google.com", "sqladmin.googleapis.com"},
		{"kmscryptokeys.kms.cnrm.cloud.google.com", "cloudkms.googleapis.com"},
		{"configconnectors.core.cnrm.cloud.google.com", ""},
		{"controllerresources.customize.core.cnrm.cloud.google.com", ""},
	}
	for _, tc := range tests {
		if got := ServiceNameForCRD(tc.crd); got != tc.want {
			t.Errorf("ServiceNameForCRD(%q) = %q, want %q", tc.crd, got, tc.want)
		}
	}
}

func TestFilterCRDNames(t *testing.T) {
	crds := []string{
		"configconnectors.core.cnrm.cloud.google.com",
		"storagebuckets.storage.cnrm.cloud.google.com",
		"sqlinstances.sql.cnrm.cloud.google.com",
		"spannerinstances.spanner.cnrm.cloud.google.com",
	}
	enabled := map[string]bool{
		"storage.custom.universe.goog":  true,
		"sqladmin.custom.universe.goog": true,
	}
	got := FilterCRDNames(crds, enabled)
	if len(got) != 3 {
		t.Fatalf("FilterCRDNames() returned %d items (%v), want 3", len(got), got)
	}
	for _, crd := range got {
		if strings.Contains(crd, "spanner") {
			t.Errorf("expected spanner CRD to be filtered out, got %v", got)
		}
	}
}

func TestDiscoverEnabledServices(t *testing.T) {
	client := &http.Client{
		Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
			if req.URL.Host != "serviceusage.custom.universe.goog" {
				t.Fatalf("unexpected host %q", req.URL.Host)
			}
			body := `{"services":[{"name":"projects/p1/services/storage.custom.universe.goog","config":{"name":"storage.custom.universe.goog"},"state":"ENABLED"},{"name":"projects/p1/services/compute.custom.universe.goog","state":"ENABLED"}]}`
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(body)),
				Header:     make(http.Header),
			}, nil
		}),
	}
	enabled, err := DiscoverEnabledServices(context.Background(), client, "p1", "custom.universe.goog")
	if err != nil {
		t.Fatalf("DiscoverEnabledServices() unexpected error: %v", err)
	}
	if !enabled["storage.custom.universe.goog"] || !enabled["compute.custom.universe.goog"] {
		t.Fatalf("unexpected enabled map: %v", enabled)
	}
}
