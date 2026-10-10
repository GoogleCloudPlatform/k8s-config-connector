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

package gcp

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

const (
	// DefaultUniverseDomain is the standard commercial Google Cloud universe domain.
	DefaultUniverseDomain = "googleapis.com"
	// UniverseDomainEnvVar is the environment variable used across Google Cloud SDKs.
	UniverseDomainEnvVar = "GOOGLE_CLOUD_UNIVERSE_DOMAIN"
)

// GetUniverseDomain returns the configured universe domain, defaulting to googleapis.com.
func GetUniverseDomain() string {
	if domain := os.Getenv(UniverseDomainEnvVar); domain != "" {
		return strings.TrimSpace(domain)
	}
	return DefaultUniverseDomain
}

// IsDefaultUniverse returns true if running against commercial googleapis.com.
func IsDefaultUniverse() bool {
	return GetUniverseDomain() == DefaultUniverseDomain
}

// FormatEndpoint constructs the fully qualified gRPC or REST endpoint for a given service.
// Examples:
//
//	FormatEndpoint("pubsub", "") -> "pubsub.googleapis.com:443" or "pubsub.custom.universe.goog:443"
//	FormatEndpoint("assuredworkloads", "europe-west3") -> "europe-west3-assuredworkloads.googleapis.com:443"
func FormatEndpoint(service string, location string) string {
	domain := GetUniverseDomain()
	if location != "" {
		return fmt.Sprintf("%s-%s.%s:443", location, service, domain)
	}
	return fmt.Sprintf("%s.%s:443", service, domain)
}

// FormatRESTURL constructs the base URL for REST-based service clients.
func FormatRESTURL(service string) string {
	return fmt.Sprintf("https://%s.%s", service, GetUniverseDomain())
}

// universeDomainRoundTripper intercepts outbound HTTP requests targeting commercial
// *.googleapis.com endpoints and rewrites them to the configured sovereign universe domain.
type universeDomainRoundTripper struct {
	base           http.RoundTripper
	universeDomain string
}

// NewUniverseDomainRoundTripper wraps an http.RoundTripper to rewrite *.googleapis.com
// hosts to *.<universeDomain> when operating in a non-default sovereign universe.
func NewUniverseDomainRoundTripper(base http.RoundTripper, universeDomain string) http.RoundTripper {
	if base == nil {
		base = http.DefaultTransport
	}
	if universeDomain == "" {
		universeDomain = GetUniverseDomain()
	}
	if universeDomain == "" || universeDomain == DefaultUniverseDomain {
		return base
	}
	return &universeDomainRoundTripper{
		base:           base,
		universeDomain: universeDomain,
	}
}

func (t *universeDomainRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.universeDomain != "" && t.universeDomain != DefaultUniverseDomain && req.URL != nil {
		if strings.Contains(req.URL.Host, ".googleapis.com") {
			clonedReq := req.Clone(req.Context())
			clonedReq.URL.Host = strings.Replace(clonedReq.URL.Host, ".googleapis.com", "."+t.universeDomain, 1)
			clonedReq.Host = clonedReq.URL.Host
			return t.base.RoundTrip(clonedReq)
		}
	}
	return t.base.RoundTrip(req)
}
