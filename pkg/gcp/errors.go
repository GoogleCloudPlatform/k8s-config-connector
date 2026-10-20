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

package gcp

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	"google.golang.org/api/googleapi"
	sqladmin "google.golang.org/api/sqladmin/v1beta4"
)

// DefaultUniverseDomain is the standard public Google Cloud universe domain.
const DefaultUniverseDomain = "googleapis.com"

var (
	// dialLookupRegex matches strings like: dial tcp: lookup <host>: no such host
	// or lookup <host> on <dns-server>: no such host
	// or lookup <host>:<port> on <dns-server>:<port>: no such host
	dialLookupRegex = regexp.MustCompile(`lookup ([a-zA-Z0-9.\-_]+(?::\d+)?)(?: on [^:]+(?::\d+)?)?: no such host`)

	// requestURLRegex matches HTTP request URLs in error strings, e.g. Get "https://...":
	requestURLRegex = regexp.MustCompile(`(?:Get|Post|Put|Delete|Patch|HEAD) "(https?://[^"]+)"`)
)

// GetConfiguredUniverseDomain returns the configured universe domain from environment variables,
// or empty string if not configured.
func GetConfiguredUniverseDomain() string {
	if ud := os.Getenv("GOOGLE_CLOUD_UNIVERSE_DOMAIN"); ud != "" {
		return strings.TrimSpace(ud)
	}
	if ud := os.Getenv("GOOGLE_UNIVERSE_DOMAIN"); ud != "" {
		return strings.TrimSpace(ud)
	}
	return ""
}

func IsNotAuthorizedError(err error) bool {
	return isGoogleErrorWithCode(err, http.StatusForbidden)
}

func IsNotFoundError(err error) bool {
	return isGoogleErrorWithCode(err, http.StatusNotFound)
}

// IsServiceEndpointUnavailableError returns true and the unresolvable host if the error
// is caused by a DNS lookup failure (NXDOMAIN / no such host) for a service endpoint
// in a non-default universe domain.
func IsServiceEndpointUnavailableError(err error) (string, bool) {
	if err == nil {
		return "", false
	}

	var reqHost string
	var lookupHost string

	// 1. Direct or wrapped *url.Error inspection
	var urlErr *url.Error
	if errors.As(err, &urlErr) {
		if urlErr.URL != "" {
			if parsed, parseErr := url.Parse(urlErr.URL); parseErr == nil {
				reqHost = parsed.Hostname()
			}
		}
	}

	// 2. Direct or wrapped *net.DNSError inspection
	var dnsErr *net.DNSError
	if errors.As(err, &dnsErr) {
		if dnsErr.IsNotFound || strings.Contains(strings.ToLower(dnsErr.Err), "no such host") {
			lookupHost = dnsErr.Name
		}
	}

	// 3. String-based matching for errors that wrap DNS lookup errors as text
	errMsg := err.Error()
	if reqHost == "" {
		if matches := requestURLRegex.FindStringSubmatch(errMsg); len(matches) > 1 {
			if parsed, parseErr := url.Parse(matches[1]); parseErr == nil {
				reqHost = parsed.Hostname()
			}
		}
	}

	if lookupHost == "" {
		if strings.Contains(errMsg, "no such host") {
			if matches := dialLookupRegex.FindStringSubmatch(errMsg); len(matches) > 1 {
				lookupHost = matches[1]
			}
		}
	}

	if lookupHost != "" {
		cleanedLookupHost := cleanHost(lookupHost)
		if isServiceHostInUniverse(cleanedLookupHost, reqHost) {
			return cleanedLookupHost, true
		}
	}

	return "", false
}

func cleanHost(host string) string {
	host = strings.TrimSpace(host)
	if h, _, err := net.SplitHostPort(host); err == nil {
		host = h
	}
	return strings.TrimSuffix(host, ".")
}

func isExcludedTarget(host string) bool {
	host = cleanHost(host)
	if host == "" {
		return true
	}
	// Localhost and loopback
	if host == "localhost" || strings.HasSuffix(host, ".localhost") || host == "127.0.0.1" || host == "::1" {
		return true
	}
	// Internal Google compute metadata server
	if host == "metadata" || host == "metadata.google.internal" || strings.HasSuffix(host, ".google.internal") {
		return true
	}
	return false
}

func isServiceHostInUniverse(lookupHost, reqHost string) bool {
	lookupHost = cleanHost(lookupHost)
	if isExcludedTarget(lookupHost) {
		return false
	}
	// If the lookup host is in the default public universe, it's not a sovereign/custom universe domain
	if lookupHost == DefaultUniverseDomain || strings.HasSuffix(lookupHost, "."+DefaultUniverseDomain) {
		return false
	}

	reqHost = cleanHost(reqHost)
	if reqHost != "" && isExcludedTarget(reqHost) {
		return false
	}

	// 1. Check against configured universe domain
	if configuredUD := GetConfiguredUniverseDomain(); configuredUD != "" && configuredUD != DefaultUniverseDomain {
		if strings.HasSuffix(lookupHost, "."+configuredUD) {
			service := strings.TrimSuffix(lookupHost, "."+configuredUD)
			if service != "" && !strings.Contains(service, ".") && isValidSubdomainLabel(service) {
				if reqHost != "" {
					reqService := extractServiceName(reqHost)
					label := strings.Split(reqHost, ".")[0]
					if reqService != "" && !strings.Contains(service, reqService) && !strings.Contains(reqService, service) &&
						!strings.Contains(service, label) && !strings.Contains(label, service) {
						return false
					}
				}
				return true
			}
		}
	}

	// 2. Check if lookupHost is in Google Cloud's canonical .universe.goog domain
	if strings.HasSuffix(lookupHost, ".universe.goog") {
		prefix := strings.TrimSuffix(lookupHost, ".universe.goog")
		// Must have at least <service>.<universe>
		if parts := strings.Split(prefix, "."); len(parts) >= 2 {
			service := parts[0]
			if isValidSubdomainLabel(service) {
				if reqHost != "" {
					reqService := extractServiceName(reqHost)
					label := strings.Split(reqHost, ".")[0]
					if reqService != "" && !strings.Contains(service, reqService) && !strings.Contains(reqService, service) &&
						!strings.Contains(service, label) && !strings.Contains(label, service) {
						return false
					}
				}
				return true
			}
		}
	}

	// 3. Check if tied to the request's service host
	if reqHost != "" {
		reqService := extractServiceName(reqHost)
		label := strings.Split(reqHost, ".")[0]
		for _, s := range []string{reqService, label} {
			if s != "" && isValidSubdomainLabel(s) && strings.HasPrefix(lookupHost, s+".") {
				domain := strings.TrimPrefix(lookupHost, s+".")
				if strings.Contains(domain, ".") && !isExcludedTarget(domain) && domain != DefaultUniverseDomain && !strings.HasSuffix(domain, "."+DefaultUniverseDomain) {
					return true
				}
			}
		}
	}

	return false
}

func extractServiceName(reqHost string) string {
	reqHost = cleanHost(reqHost)
	if reqHost == "" {
		return ""
	}
	parts := strings.Split(reqHost, ".")
	if len(parts) == 0 {
		return ""
	}
	label := parts[0]
	// Handle regional endpoints like us-central1-aiplatform -> aiplatform
	if strings.Contains(label, "-") {
		subParts := strings.Split(label, "-")
		return subParts[len(subParts)-1]
	}
	return label
}

func isValidSubdomainLabel(s string) bool {
	if len(s) == 0 || len(s) > 63 {
		return false
	}
	for i, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			continue
		}
		if r == '-' && i > 0 && i < len(s)-1 {
			continue
		}
		return false
	}
	return true
}

func isGoogleErrorWithCode(err error, code int) bool {
	if err == nil {
		return false
	}
	ge := &googleapi.Error{}
	if errors.As(err, &ge) {
		return ge.Code == code
	}
	return false
}

func FormatSQLOperationErrorMessages(opErrs []*sqladmin.OperationError) string {
	var errMessages []string
	for _, opErr := range opErrs {
		errMessages = append(errMessages, fmt.Sprintf("%v %v", opErr.Code, opErr.Message))
	}
	return strings.Join(errMessages, ", ")
}
