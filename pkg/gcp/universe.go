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
	"net/http"
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

// RewriteHostForUniverse rewrites a request host and path for a custom universe domain.
// When the host targets www.googleapis.com or www.<universeDomain> with a service path
// such as "/compute/beta/...", it routes the host to "<service>.<universeDomain>".
func RewriteHostForUniverse(host, urlPath, universeDomain string) string {
	domain := ResolveUniverseDomain(universeDomain)
	if !IsCustomUniverse(domain) || host == "" {
		return host
	}

	hostname := host
	port := ""
	if idx := strings.LastIndex(host, ":"); idx != -1 {
		hostname = host[:idx]
		port = host[idx:]
	}

	if strings.EqualFold(hostname, DefaultUniverseDomain) {
		hostname = domain
	} else if strings.HasSuffix(strings.ToLower(hostname), "."+DefaultUniverseDomain) {
		prefix := hostname[:len(hostname)-len("."+DefaultUniverseDomain)]
		hostname = prefix + "." + domain
	}

	if strings.EqualFold(hostname, "www."+domain) {
		if svc := extractServiceFromWWWPath(urlPath); svc != "" {
			hostname = svc + "." + domain
		}
	}

	return hostname + port
}

func extractServiceFromWWWPath(urlPath string) string {
	trimmed := strings.TrimPrefix(urlPath, "/")
	if trimmed == "" {
		return ""
	}
	parts := strings.Split(trimmed, "/")
	first := parts[0]
	if first == "upload" || first == "resumable" {
		if len(parts) > 1 {
			first = parts[1]
		} else {
			return ""
		}
	}
	if first == "" || first == "oauth2" || first == " batch" || first == "batch" {
		return ""
	}
	return first
}

// UniverseDomainRoundTripper rewrites outgoing HTTP requests targeting *.googleapis.com
// to the configured custom universe domain, including routing www.googleapis.com/<service>/...
// requests to <service>.<universeDomain>.
type UniverseDomainRoundTripper struct {
	Inner          http.RoundTripper
	UniverseDomain string
}

// NewUniverseDomainRoundTripper wraps inner with UniverseDomainRoundTripper when a custom
// universe domain is configured.
func NewUniverseDomainRoundTripper(inner http.RoundTripper, universeDomain string) http.RoundTripper {
	if inner == nil {
		inner = http.DefaultTransport
	}
	if existing, ok := inner.(*UniverseDomainRoundTripper); ok {
		if universeDomain != "" {
			existing.UniverseDomain = universeDomain
		}
		return existing
	}
	return &UniverseDomainRoundTripper{
		Inner:          inner,
		UniverseDomain: universeDomain,
	}
}

func (rt *UniverseDomainRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	inner := rt.Inner
	if inner == nil {
		inner = http.DefaultTransport
	}
	domain := ResolveUniverseDomain(rt.UniverseDomain)
	if !IsCustomUniverse(domain) || req == nil || req.URL == nil {
		return inner.RoundTrip(req)
	}

	cloned := req.Clone(req.Context())
	newHost := RewriteHostForUniverse(cloned.URL.Host, cloned.URL.Path, domain)
	if newHost != cloned.URL.Host {
		cloned.URL.Host = newHost
		cloned.Host = newHost
	}
	return inner.RoundTrip(cloned)
}
