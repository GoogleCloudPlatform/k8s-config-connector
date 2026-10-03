// Copyright 2024 Google LLC
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
	"net/http"
	"os"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp"
)

func TestFormatServiceAccountEmail(t *testing.T) {
	testCases := []struct {
		Name           string
		ProjectID      string
		SAName         string
		ExpectedResult string
	}{
		{
			Name:           "Standard commercial project",
			ProjectID:      "my-standard-project",
			SAName:         "kcc-system",
			ExpectedResult: "kcc-system@my-standard-project.iam.gserviceaccount.com",
		},
		{
			Name:           "Sovereign Germany dogfood project (eu0 partition)",
			ProjectID:      "eu0:kcc-eval-de",
			SAName:         "kcc-system",
			ExpectedResult: "kcc-system@kcc-eval-de.eu0.iam.gserviceaccount.com",
		},
		{
			Name:           "Sovereign France project (fr0 partition)",
			ProjectID:      "fr0:sovereign-production",
			SAName:         "kcc-system",
			ExpectedResult: "kcc-system@sovereign-production.fr0.iam.gserviceaccount.com",
		},
		{
			Name:           "Custom air-gapped partition (sec1 partition)",
			ProjectID:      "sec1:classified-workload",
			SAName:         "app-agent",
			ExpectedResult: "app-agent@classified-workload.sec1.iam.gserviceaccount.com",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := gcp.FormatServiceAccountEmail(tc.ProjectID, tc.SAName)
			if result != tc.ExpectedResult {
				t.Fatalf("FormatServiceAccountEmail(%q, %q) = %q, want %q", tc.ProjectID, tc.SAName, result, tc.ExpectedResult)
			}
		})
	}
}

func TestFormatWorkloadPool(t *testing.T) {
	testCases := []struct {
		Name           string
		ProjectID      string
		ExpectedResult string
	}{
		{
			Name:           "Standard commercial project",
			ProjectID:      "my-standard-project",
			ExpectedResult: "my-standard-project.svc.id.goog",
		},
		{
			Name:           "Sovereign Germany dogfood project",
			ProjectID:      "eu0:kcc-eval-de",
			ExpectedResult: "kcc-eval-de.eu0.svc.id.goog",
		},
		{
			Name:           "Sovereign France project",
			ProjectID:      "fr0:sovereign-production",
			ExpectedResult: "sovereign-production.fr0.svc.id.goog",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := gcp.FormatWorkloadPool(tc.ProjectID)
			if result != tc.ExpectedResult {
				t.Fatalf("FormatWorkloadPool(%q) = %q, want %q", tc.ProjectID, result, tc.ExpectedResult)
			}
		})
	}
}

func TestFormatEndpoint(t *testing.T) {
	orig := os.Getenv(gcp.UniverseDomainEnvVar)
	defer os.Setenv(gcp.UniverseDomainEnvVar, orig)

	// Test default universe
	os.Unsetenv(gcp.UniverseDomainEnvVar)
	if ep := gcp.FormatEndpoint("pubsub", ""); ep != "pubsub.googleapis.com:443" {
		t.Errorf("expected pubsub.googleapis.com:443, got %q", ep)
	}
	if ep := gcp.FormatEndpoint("assuredworkloads", "europe-west3"); ep != "europe-west3-assuredworkloads.googleapis.com:443" {
		t.Errorf("expected europe-west3-assuredworkloads.googleapis.com:443, got %q", ep)
	}

	// Test sovereign universe (Germany)
	os.Setenv(gcp.UniverseDomainEnvVar, "apis-berlin-build0.goog")
	if ep := gcp.FormatEndpoint("pubsub", ""); ep != "pubsub.apis-berlin-build0.goog:443" {
		t.Errorf("expected pubsub.apis-berlin-build0.goog:443, got %q", ep)
	}
	if ep := gcp.FormatEndpoint("assuredworkloads", "u-germany-northeast1"); ep != "u-germany-northeast1-assuredworkloads.apis-berlin-build0.goog:443" {
		t.Errorf("expected u-germany-northeast1-assuredworkloads.apis-berlin-build0.goog:443, got %q", ep)
	}

	// Test sovereign universe (France)
	os.Setenv(gcp.UniverseDomainEnvVar, "apis-paris-build0.goog")
	if ep := gcp.FormatEndpoint("pubsub", ""); ep != "pubsub.apis-paris-build0.goog:443" {
		t.Errorf("expected pubsub.apis-paris-build0.goog:443, got %q", ep)
	}
}

type mockRoundTripper struct {
	lastReq *http.Request
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	m.lastReq = req
	return &http.Response{StatusCode: 200}, nil
}

func TestUniverseDomainRoundTripper(t *testing.T) {
	mock := &mockRoundTripper{}
	rt := gcp.NewUniverseDomainRoundTripper(mock, "apis-berlin-build0.goog")

	req, err := http.NewRequest("GET", "https://iam.googleapis.com/v1/projects/eu0:kcc-eval-de/serviceAccounts/kcc-test-sa@eu0:kcc-eval-de.iam.gserviceaccount.com", nil)
	if err != nil {
		t.Fatalf("unexpected error creating request: %v", err)
	}

	_, err = rt.RoundTrip(req)
	if err != nil {
		t.Fatalf("unexpected error in RoundTrip: %v", err)
	}

	if mock.lastReq.URL.Host != "iam.apis-berlin-build0.goog" {
		t.Errorf("expected host iam.apis-berlin-build0.goog, got %q", mock.lastReq.URL.Host)
	}
	expectedPath := "/v1/projects/eu0:kcc-eval-de/serviceAccounts/kcc-test-sa@kcc-eval-de.eu0.iam.gserviceaccount.com"
	if mock.lastReq.URL.Path != expectedPath {
		t.Errorf("expected path %q, got %q", expectedPath, mock.lastReq.URL.Path)
	}

	// Test www.googleapis.com mapping to service subdomain
	wwwReq, err := http.NewRequest("GET", "https://www.googleapis.com/compute/beta/projects/eu0:kcc-eval-de/zones/u-germany-northeast1-a/instanceGroupManagers", nil)
	if err != nil {
		t.Fatalf("unexpected error creating www request: %v", err)
	}
	_, err = rt.RoundTrip(wwwReq)
	if err != nil {
		t.Fatalf("unexpected error in RoundTrip: %v", err)
	}
	if mock.lastReq.URL.Host != "compute.apis-berlin-build0.goog" {
		t.Errorf("expected host compute.apis-berlin-build0.goog, got %q", mock.lastReq.URL.Host)
	}
}
