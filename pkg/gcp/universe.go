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
	"os"
	"strings"
)

const (
	// DefaultUniverseDomain is the standard commercial Google Cloud universe domain.
	DefaultUniverseDomain = "googleapis.com"
)

// UniverseDomainEnvVars lists the environment variables checked in priority order
// when resolving the active Google Cloud universe domain.
var UniverseDomainEnvVars = []string{
	"GOOGLE_CLOUD_UNIVERSE_DOMAIN",
	"UNIVERSE_DOMAIN",
}

// ResolveUniverseDomain returns the configured universe domain, checking explicit
// configuration first, then environment variables, and defaulting to "googleapis.com".
func ResolveUniverseDomain(explicit string) string {
	trimmed := strings.TrimSpace(explicit)
	if trimmed != "" {
		return strings.TrimPrefix(trimmed, ".")
	}
	for _, envKey := range UniverseDomainEnvVars {
		if val := strings.TrimSpace(os.Getenv(envKey)); val != "" {
			return strings.TrimPrefix(val, ".")
		}
	}
	return DefaultUniverseDomain
}

// IsCustomUniverse returns true when the given universe domain is non-empty
// and differs from the commercial default ("googleapis.com").
func IsCustomUniverse(universeDomain string) bool {
	resolved := strings.TrimPrefix(strings.TrimSpace(universeDomain), ".")
	return resolved != "" && !strings.EqualFold(resolved, DefaultUniverseDomain)
}

// ServiceHost returns "<service>.<universeDomain>" for the given GCP service prefix.
func ServiceHost(service, universeDomain string) string {
	domain := ResolveUniverseDomain(universeDomain)
	svc := strings.TrimSuffix(strings.TrimSpace(service), "."+DefaultUniverseDomain)
	svc = strings.TrimSuffix(svc, "."+domain)
	return svc + "." + domain
}

// ReplaceUniverseDomain replaces "googleapis.com" in an endpoint or URL with the
// resolved universe domain when a custom universe is active.
func ReplaceUniverseDomain(endpoint, universeDomain string) string {
	domain := ResolveUniverseDomain(universeDomain)
	if !IsCustomUniverse(domain) {
		return endpoint
	}
	return strings.ReplaceAll(endpoint, DefaultUniverseDomain, domain)
}
