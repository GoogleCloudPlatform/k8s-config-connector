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

package config

import "testing"

func TestGetUniverseDomain(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		env      string
		expected string
	}{
		{
			name:     "unset_defaultsToPublicUniverse",
			expected: DefaultUniverseDomain,
		},
		{
			name:     "fieldSet_usesField",
			field:    "example-apis.test",
			expected: "example-apis.test",
		},
		{
			name:     "envSet_usesEnv",
			env:      "example-apis.test",
			expected: "example-apis.test",
		},
		{
			name:     "fieldAndEnvSet_fieldWins",
			field:    "field-apis.test",
			env:      "env-apis.test",
			expected: "field-apis.test",
		},
		{
			name:     "fieldSetToPublicUniverse_isHonoured",
			field:    DefaultUniverseDomain,
			env:      "env-apis.test",
			expected: DefaultUniverseDomain,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.env != "" {
				t.Setenv(UniverseDomainEnvVar, tc.env)
			}
			c := &ControllerConfig{UniverseDomain: tc.field}
			if got := c.GetUniverseDomain(); got != tc.expected {
				t.Errorf("GetUniverseDomain() = %q, want %q", got, tc.expected)
			}
		})
	}
}

func TestGetUniversePrefix(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		env      string
		expected string
	}{
		{
			name:     "unset_isEmpty",
			expected: "",
		},
		{
			name:     "fieldSet_usesField",
			field:    "zzz",
			expected: "zzz",
		},
		{
			name:     "envSet_usesEnv",
			env:      "zzz",
			expected: "zzz",
		},
		{
			name:     "fieldAndEnvSet_fieldWins",
			field:    "field-prefix",
			env:      "env-prefix",
			expected: "field-prefix",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.env != "" {
				t.Setenv(UniversePrefixEnvVar, tc.env)
			}
			c := &ControllerConfig{UniversePrefix: tc.field}
			if got := c.GetUniversePrefix(); got != tc.expected {
				t.Errorf("GetUniversePrefix() = %q, want %q", got, tc.expected)
			}
		})
	}
}

func TestIsDefaultUniverse(t *testing.T) {
	tests := []struct {
		name     string
		domain   string
		expected bool
	}{
		{name: "empty_isDefault", domain: "", expected: true},
		{name: "explicitPublicUniverse_isDefault", domain: DefaultUniverseDomain, expected: true},
		{name: "otherUniverse_isNotDefault", domain: "example-apis.test", expected: false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			c := &ControllerConfig{UniverseDomain: tc.domain}
			if got := c.IsDefaultUniverse(); got != tc.expected {
				t.Errorf("IsDefaultUniverse() = %v, want %v", got, tc.expected)
			}
		})
	}
}

func TestEndpoint(t *testing.T) {
	const universe = "example-apis.test"

	tests := []struct {
		name     string
		domain   string
		endpoint string
		expected string
	}{
		{
			name:     "publicUniverse_returnsEndpointUnchanged",
			domain:   "",
			endpoint: "networksecurity.googleapis.com:443",
			expected: "networksecurity.googleapis.com:443",
		},
		{
			name:     "hostAndPort",
			domain:   universe,
			endpoint: "networksecurity.googleapis.com:443",
			expected: "networksecurity." + universe + ":443",
		},
		{
			name:     "regionalPrefix",
			domain:   universe,
			endpoint: "us-central1-aiplatform.googleapis.com:443",
			expected: "us-central1-aiplatform." + universe + ":443",
		},
		{
			name:     "regionalEndpointForm",
			domain:   universe,
			endpoint: "parametermanager.us-central1.rep.googleapis.com:443",
			expected: "parametermanager.us-central1.rep." + universe + ":443",
		},
		{
			name:     "schemePrefixed",
			domain:   universe,
			endpoint: "https://us-central1-aiplatform.googleapis.com",
			expected: "https://us-central1-aiplatform." + universe,
		},
		{
			name:     "previewPrefixed",
			domain:   universe,
			endpoint: "public-preview-recaptchaenterprise.googleapis.com:443",
			expected: "public-preview-recaptchaenterprise." + universe + ":443",
		},
		{
			name:     "hostOnly",
			domain:   universe,
			endpoint: "us-central1-cloudresourcemanager.googleapis.com",
			expected: "us-central1-cloudresourcemanager." + universe,
		},
		{
			name:     "endpointWithoutPublicSuffix_isUnchanged",
			domain:   universe,
			endpoint: "localhost:8080",
			expected: "localhost:8080",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			c := &ControllerConfig{UniverseDomain: tc.domain}
			if got := c.Endpoint(tc.endpoint); got != tc.expected {
				t.Errorf("Endpoint(%q) = %q, want %q", tc.endpoint, got, tc.expected)
			}
		})
	}
}

// TestUniverseDomainAndPrefixAreIndependent guards the invariant that the
// universe domain and the universe prefix are unrelated values.
//
// It is tempting to derive one from the other, because in the universes we know
// about they share a stem (domain "s3nsapis.fr" with prefix "s3ns"). That is a
// coincidence, not a rule. This test uses a pair with no shared stem, so any
// future attempt to compute one from the other fails here.
func TestUniverseDomainAndPrefixAreIndependent(t *testing.T) {
	c := &ControllerConfig{
		UniverseDomain: "example-apis.test",
		UniversePrefix: "zzz",
	}

	if got, want := c.GetUniverseDomain(), "example-apis.test"; got != want {
		t.Errorf("GetUniverseDomain() = %q, want %q", got, want)
	}
	if got, want := c.GetUniversePrefix(), "zzz"; got != want {
		t.Errorf("GetUniversePrefix() = %q, want %q", got, want)
	}
	if got, want := c.Endpoint("compute.googleapis.com:443"), "compute.example-apis.test:443"; got != want {
		t.Errorf("Endpoint() = %q, want %q", got, want)
	}
}
